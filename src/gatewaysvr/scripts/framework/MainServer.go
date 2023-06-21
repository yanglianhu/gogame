/*package framework

import (
	"encoding/binary"
	"fmt"
	"gatewaysvr/scripts/Apps"
	"gatewaysvr/scripts/ConnectionMgr"
	"gatewaysvr/scripts/config"
	"lgameserver/pkg/cs"
	"lgameserver/pkg/ss"
	"lgameserver/src/base/logger"
	"lgameserver/src/base/relay"
	"lgameserver/src/base/utils"
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/walkon/gnet"
	gerrors "github.com/walkon/gnet/errors"
	"x11-gitlab.diezhi.net/laide/utils/snowflake"
)

var gamesvrConnected bool

type ServerState int

const (
	ServerState_Init ServerState = 0
	// 正常运行阶段
	ServerState_Normal ServerState = 1
	// 准备退出阶段，需要做停服退出准备，不再处理客户端的请求
	ServerState_PreStop ServerState = 2
	// 退出
	ServerState_Stop ServerState = 3
)

type MainServer struct {
	*gnet.EventServer
	svr          gnet.Server
	network      string
	addr         string
	multicore    bool
	async        bool
	started      int32
	connected    int32
	clientActive int32
	disconnected int32

	// workerPool   *goroutine.Pool
	codec gnet.ICodec

	state ServerState

	// 人为制造时间延迟，单位: 毫秒
	latency int
	rateObj *utils.RateLimiter
}

func (this *MainServer) OnInitComplete(svr gnet.Server) (action gnet.Action) {
	rand.Seed(time.Now().UnixNano())

	this.svr = svr
	this.state = ServerState_Normal
	utils.AddEventHandler(ss.EventType_EId_StopService, this.handleStopServiceEvent)
	utils.AddEventHandler(ss.EventType_EId_Latency, this.latencyEvent)
	utils.AddTimer("MainServer::tick1s", time.Second, this.tick1s)
	relay.RelayMgr.Init(config.GetServiceId(), int32(ss.ServiceType_GatewaySvr), 0, &svr, logger.Log)
	return
}

func (this *MainServer) handleStopServiceEvent(event *utils.Event) {
	this.state = ServerState_PreStop
	logger.Log.Error("service stoping ...")
}

func (this *MainServer) latencyEvent(event *utils.Event) {
	if event == nil || len(event.Param) != 1 {
		return
	}
	this.latency = int(event.Param[0])
}

func (this *MainServer) PollerPreInit() {
}

func (this *MainServer) isStop() bool {
	return (this.state == ServerState_PreStop || this.state == ServerState_Stop)
}

func (this *MainServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	if this.isStop() {
		action = gnet.Close
		logger.Log.Error("[OnOpened] server stoping, cann't process anything, state: %d", this.state)
		return
	}
	// 服务之间的连接
	if c.Context() != nil {
		// service connect success
		logger.Log.Error("service connect success. connIdx: %#v", c.Context())
		return
	}
	connIdx := snowflake.GenUniqId(snowflake.IDType_ConnIdx)
	c.SetContext(connIdx)
	ConnectionMgr.ConnMgr.AddGnetConn(connIdx, &c)
	atomic.AddInt32(&this.connected, 1)

	return
}

func (this *MainServer) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	if this.isStop() {
		logger.Log.Error("[OnClosed] server stoping, cann't process anything")
		return
	}
	connidx, ok := c.Context().(uint64)
	if ok {
		gamesvrConnId := utils.ServiceConnIdx(ss.ServiceType_GameSvr, 1)
		if connidx == gamesvrConnId {
			gamesvrConnected = false
		} else {
			//从客户端发来的断开连接
			if connidx >= utils.MaxServiceConnIdx {
				userConn := ConnectionMgr.ConnMgr.GetUserConnByConnId(connidx)
				if userConn != nil {
					var ssHead *ss.SSHead = &ss.SSHead{}
					ssHead.CmdId = uint32(ss.EnmCmdValue_ECV_SSPlayerLogout)
					ssHead.AccId = userConn.AccId
					var ssMsg *ss.SSPlayerLogout = &ss.SSPlayerLogout{AccId: ssHead.AccId, Reason: "user disconnect"}
					utils.SendMsg(utils.GamesvrConn, ssHead, ssMsg)
					Apps.SendOffline2OnlineSvr(userConn.AccId)
					ConnectionMgr.ConnMgr.RemoveConnection(userConn)
				}

			}
		}
	}
	return
}

func (this *MainServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if this.isStop() {
		action = gnet.Close
		logger.Log.Error("[React] server stoping, cann't process anything")
		return
	}

	connIdx, ok := c.Context().(uint64)
	if ok && connIdx < utils.MaxServiceConnIdx {
		return this.reactService(connIdx, frame, c)
	}

	return this.reactClient(connIdx, frame, c)
}

func (mainServer *MainServer) reactClient(connIdx uint64, frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	now := time.Now()
	userConn := ConnectionMgr.ConnMgr.GetUserConnByConnId(connIdx)
	if userConn == nil {
		logger.Log.Error("userConn is nil")
		return
	}
	if connIdx > 0 && !mainServer.rateObj.GetAllowed(fmt.Sprintf("%d", connIdx), now.Unix()) {
		action = gnet.Close
		logger.Log.Error("too many packets")
		return
	}

	var length = uint32(len(frame))
	if length <= 2 || length < (2+uint32(frame[1])) {
		action = gnet.Close
		logger.Log.Error("length too small. length: %d", length)
		return
	}

	var (
		xxteaKey []byte
		buffer   []byte
	)

	// 加密 && 压缩
	bitmask := uint32(frame[0])
	// 头部长度
	headlen := uint32(frame[1])

	// 加密
	if (bitmask & 0xf) != 0 {
		if userConn.EncryptKey == nil {
			xxteaKey = []byte(config.Conf.Common.XXTeaKey)
		} else {
			xxteaKey = userConn.EncryptKey
		}
		buffer = utils.XXTeaDecryptAndDecompress(frame[2:], xxteaKey)
		if buffer == nil {
			logger.Log.Error("buff Decrypt is nil,key==%s", string(xxteaKey))
			return
		}
	} else {
		buffer = frame[2:]
	}

	head := &cs.CSHead{}

	if len(buffer) < int(headlen) {
		action = gnet.Close
		logger.Log.Error("buffer is too small. buf.len: %d, headlen: %d, frame: %#v", len(buffer), headlen, frame)
		return
	}

	if err := head.Unmarshal(buffer[:headlen]); err != nil {
		action = gnet.Close
		logger.Log.Error("unpack head failed. length: %d, headlen: %d", length, headlen)
		return
	}
	if head.SeqId != userConn.LastSeqId+1 {
		action = gnet.Close
		logger.Log.Error("head.SeqId too small:%d", head.SeqId)
		return
	}

	defer func() {
		utils.StatMgr.StatCmdCall(head.GetCmdId(), len(frame), time.Since(now))
	}()

	imessage := cs.NewClientMessage(int32(head.GetCmdId()))
	if imessage == nil {
		action = gnet.Close
		logger.Log.Error("not found message. CmdId: %d", head.GetCmdId())
		return
	}

	moffset := headlen

	if merr := imessage.Unmarshal(buffer[moffset:]); merr != nil {
		action = gnet.Close
		logger.Log.Error("unpack message failed. Head: %s, Err: %#v", head.String(), merr)
		return
	}

	if config.Conf.Common.NetLog {
		logger.Log.Debug("NetLog Recv Message(client). Head: %s, Message: %s", head.String(), imessage.String())
	}
	accId := userConn.AccId
	if accId == 0 && head.GetCmdId() != uint32(cs.EnmCmdValue_ECV_CSLoginReq) {
		action = gnet.Close
		logger.Log.Error("send msg before login %d", head.GetCmdId())
		return
	}
	if head.GetCmdId() == uint32(cs.EnmCmdValue_ECV_CSLoginReq) {
		req, ok := imessage.(*cs.CSLoginReq)
		if ok && req.UseEncrypt {
			sKey, _ := utils.GenerateRandomString()
			userConn.EncryptKey = []byte(sKey)
		}
	}
	head.Context = connIdx
	userConn.LastSeqId = head.SeqId
	utils.SendMsgFromClient(accId, connIdx, head, imessage)
	//go msgqueue.MsgQueueMgr.AddCSMsg2Queue(connIdx, player.GetAccId(), head, imessage)
	return
}

func (mainServer *MainServer) reactService(connIdx uint64, frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	var length = uint32(len(frame))
	if length <= 2 || length < (2+uint32(frame[1])) {
		action = gnet.Close
		logger.Log.Error("%d|length too small. length: %d", connIdx, length)
		return
	}

	// 加密 && 压缩
	// bit := uint32(frame[0])
	// 头部长度
	headlen := uint32(frame[1])
	head := &ss.SSHead{}

	if err := head.Unmarshal(frame[2 : 2+headlen]); err != nil {
		action = gnet.Close
		logger.Log.Error("%d|unpack head failed. length: %d, headlen: %d", connIdx, length, headlen)
		return
	}
	now := time.Now()
	defer func() {
		utils.StatMgr.StatCmdCall(head.GetCmdId(), len(frame), time.Since(now))
	}()
	moffset := 2 + headlen
	if head.GetCmdId() != uint32(ss.EnmCmdValue_ECV_SSTransReq) {
		imessage := ss.NewServiceMessage(int32(head.GetCmdId()))
		if imessage == nil {
			action = gnet.Close
			logger.Log.Error("%d|not found message. CmdId: %d", connIdx, head.GetCmdId())
			return
		}

		if merr := imessage.Unmarshal(frame[moffset:]); merr != nil {
			action = gnet.Close
			logger.Log.Error("%d|unpack message failed. Head: %s, Err: %#v", connIdx, head.String(), merr)
			return
		}
		if head.GetCmdId() == uint32(ss.EnmCmdValue_ECV_SSSend2Client) {
			Apps.DealMsgFromGamesvr2Client(head, imessage)
		} else {

			handler := Apps.GetServiceHandle(head.GetCmdId())
			if handler == nil {
				logger.Log.Error("not found handler. Head: %s, Message: %s", head.String(), imessage.String())
				return
			}
			handler(c, head, imessage)
		}
	} else {
		Apps.HandleTransReq(head, frame[moffset:])
	}
	return
}

func (this *MainServer) tick1s(now time.Time) {
	for !gamesvrConnected {
		logger.Log.Debug("gamesvrConnected == %v", gamesvrConnected)
		gamesvrConnected = newGamesvrConn(config.Conf.Network.GamesvrUrl, &this.svr)
	}
	ConnectionMgr.ConnMgr.KickIdleConnection()
}

func (this *MainServer) Tick() (delay time.Duration, action gnet.Action) {
	// @NOTE: 不能在这里添加代码, 这里是子线程
	delay = time.Millisecond * 1000
	return
}

// 返回: errors.ErrServerShutdown 则关闭服务器
func (this *MainServer) PollerProc() error {
	utils.ResetTime()

	now := time.Now()

	utils.WorkerPool.Callback()
	utils.ProcessTimer()

	// 服务器退出
	if this.state == ServerState_Stop {
		logger.Log.Error("server Stop ...")
		// action = gnet.Shutdown
		return gerrors.ErrServerShutdown
	}

	utils.StatMgr.StatFuncCall("PollerProc", time.Since(now))

	return nil
}

func newGamesvrConn(gamesvrUrl string, svr *gnet.Server) bool {
	connFd, err := gnet.NewTCPConnFd("tcp", gamesvrUrl, gnet.WithTCPNoDelay(gnet.TCPDelay))
	if err != nil {
		logger.Log.Error("connect service failed. Host: %s, Err: %#v", gamesvrUrl, err)
		return false
	} else {
		ConnIndex := utils.ServiceConnIdx(ss.ServiceType_GameSvr, 1)
		pconn, cerr := gnet.AddTCPConnector(svr, connFd, ConnIndex)
		if cerr != nil {
			logger.Log.Error("add service to poller failed. Host: %s, Err: %#v", gamesvrUrl, cerr)
			return false
		} else {
			conn := gnet.Conn(pconn)
			conn.SetContext(ConnIndex)
			utils.GamesvrConn = conn
			logger.Log.Error("connect %s success,ConnIndex:%#v", gamesvrUrl, ConnIndex)
			return true
		}
	}
}

func StartServe() {
	reuseport := true
	// multicore 为 false，否则有可能会造成主逻辑是多线程
	multicore := false
	async := true
	network := "tcp"
	addr := config.Conf.Network.ServerService

	encoderConfig := gnet.EncoderConfig{
		ByteOrder:                       binary.BigEndian,
		LengthFieldLength:               4,
		LengthAdjustment:                0,
		LengthIncludesLengthFieldLength: false,
	}

	decoderConfig := gnet.DecoderConfig{
		ByteOrder:           binary.BigEndian,
		LengthFieldOffset:   0,
		LengthFieldLength:   4,
		LengthAdjustment:    0,
		InitialBytesToStrip: 4,
	}

	codec := gnet.NewLengthFieldBasedFrameCodec(encoderConfig, decoderConfig)

	ms := &MainServer{
		network:   network,
		addr:      addr,
		multicore: multicore,
		async:     async,
		// workerPool: goroutine.Default(),
		codec: codec,
	}

	logger.Log.Error("start server service ...")
	var MaxPackPerMinute uint32
	if config.Conf.Common.MaxPackPerMinute > 0 {
		MaxPackPerMinute = config.Conf.Common.MaxPackPerMinute
	} else {
		MaxPackPerMinute = 200
	}
	ms.rateObj = utils.NewRateObj(MaxPackPerMinute)
	err := gnet.Serve(ms, network+"://"+addr,
		gnet.WithLockOSThread(async),
		gnet.WithMulticore(multicore),
		gnet.WithReusePort(reuseport),
		gnet.WithTicker(true),
		gnet.WithTCPKeepAlive(time.Minute*1),
		gnet.WithCodec(codec),
		gnet.WithTCPNoDelay(gnet.TCPDelay),
		gnet.WithLoadBalancing(gnet.SourceAddrHash),
		gnet.WithSocketRecvBuffer(1024*20),
		// gnet.WithLogger(logger.Log),
	)
	if err != nil {
		logger.Log.Error("gnet.Serve err:%v", err)
	}

}*/
