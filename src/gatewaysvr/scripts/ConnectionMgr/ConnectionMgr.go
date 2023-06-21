package ConnectionMgr

import (
	"encoding/binary"
	"errors"
	"fmt"
	"gatewaysvr/scripts/CommDef"
	"net"
	"strings"
	"sync"
	"time"
)

func parseClientMsg(data []byte) (*cs.CSHead, cs.IMessage, error) {
	if len(data) < 4 {
		return nil, nil, errors.New("kcp data error")
	}
	length := binary.BigEndian.Uint32(data[:4])
	if len(data) != int(4+length) {
		return nil, nil, errors.New("kcp data error")
	}
	//pack := make([]byte, int(4+length))
	headLen := int32(data[5])
	headBuff := data[6 : 6+headLen]
	MsgBuff := data[6+headLen:]
	var head cs.CSHead
	head.Unmarshal(headBuff)
	imessage := cs.NewClientMessage(int32(head.GetCmdId()))
	if imessage == nil {
		return nil, nil, nil
	}
	imessage.Unmarshal(MsgBuff)
	return &head, imessage, nil
}

func handleConnection(userConn *CommDef.UserConnector) {
	// 握手验证
	data := make([]byte, 1024)
	//等待3秒钟，等待发握手连接
	userConn.KcpConn.SetReadDeadline(time.Now().Add(time.Second * 3))
	_, err := userConn.KcpConn.Read(data)
	if err != nil {
		logger.Log.Error("read error:%s", err.Error())
		userConn.KcpConn.Close()
		return
	}
	//第一个包必须是握手包
	head, _, err := parseClientMsg(data)
	if err != nil {
		logger.Log.Error("parseClientMsg err :%s", err.Error())
		userConn.KcpConn.Close()
		return
	}
	if head.CmdId != uint32(cs.EnmCmdValue_ECV_CSLoginReq) {
		logger.Log.Error("send data before login")
		userConn.KcpConn.Close()
		return
	}
	// 添加到连接管理器
	ConnMgr.AddConnection(userConn)

	// 读取客户端数据并转发到游戏服务器
	for {
		data := make([]byte, 1024)
		userConn.KcpConn.SetReadDeadline(time.Now().Add(time.Second * 3))
		n, err := userConn.KcpConn.Read(data)
		if err != nil {
			if strings.Contains(err.Error(), "timeout") {
				time.Sleep(time.Millisecond * 20)
				continue
			} else {
				break
			}
		}
		fmt.Println("n==%d", n)

		// 转发数据给游戏服务器
		//gameServer.Send(data[:n])
	}

	// 从连接管理器中移除连接
	userConn.KcpConn.Close()
	ConnMgr.RemoveConnection(userConn)
	//fmt.Println("connection closed")
}

// 连接管理器，用于管理客户端连接
type ConnectionManager struct {
	AllConn    sync.Map
	AccId2Conn sync.Map
}

// 添加连接
func (m *ConnectionManager) AddConnection(conn *CommDef.UserConnector) {
	m.AllConn.Store(conn, true)
}

// 添加连接
func (m *ConnectionManager) BindUid(conn *CommDef.UserConnector, Uid uint64) {
	logger.Log.Debug("BindAccId----------%d", Uid)
	conn.AccId = Uid
	conn.Status = CommDef.ConnectionStat_Authed
	m.AccId2Conn.Store(AccId, conn)
}

func (m *ConnectionManager) GetUserConnByAccid(Uid uint64) *CommDef.UserConnector {
	ret, ok := m.AccId2Conn.Load(Uid)
	if ok {
		return ret.(*CommDef.UserConnector)
	}
	return nil
}

// 移除连接
func (m *ConnectionManager) RemoveConnection(conn *CommDef.UserConnector) {
	if _, ok := m.AllConn.Load(conn.Context); ok {
		m.AllConn.Delete(conn.Context)
		m.AccId2Conn.Delete(conn.AccId)
	}
}

func (m *ConnectionManager) AddTcpConn(connIdx uint64, conn *net.Conn) {
	logger.Log.Debug("AddGnetConn----------%d", connIdx)
	userConn := newUserGnetConn(conn)
	m.AllConn.Store(connIdx, userConn)
}

func (m *ConnectionManager) GetUserConnByConnId(connIdx uint64) *CommDef.UserConnector {
	ret, ok := m.AllConn.Load(connIdx)
	if ok {
		return ret.(*CommDef.UserConnector)
	}
	return nil
}

// 踢掉空连接
func (m *ConnectionManager) KickIdleConnection() {
	count := 0
	m.AllConn.Range(func(key, value interface{}) bool {
		UserConnector, ok := value.(*CommDef.UserConnector)
		if ok {
			if time.Now().UnixMilli()-UserConnector.CreateTime > 3000 && UserConnector.Status == 0 {
				(*UserConnector.TcpConn).Close()
				m.RemoveConnection(UserConnector)
				count++
				if count > 1000 {
					return false
				}
			}
		}
		return true
	})
}

// 广播消息给所有连接
/*func (m *ConnectionManager) Broadcast(data []byte) {
	for conn := range m.connections {
		conn.Write(data)
	}
}*/

// 初始化连接管理器和游戏服务器
var ConnMgr = &ConnectionManager{}

func newUserGnetConn(conn *net.Conn) *CommDef.UserConnector {
	userConn := &CommDef.UserConnector{}
	userConn.CreateTime = time.Now().UnixMilli()
	userConn.TcpConn = conn
	return userConn
}
