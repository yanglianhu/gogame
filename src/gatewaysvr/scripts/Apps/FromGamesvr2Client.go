package Apps

/*import (
	"gatewaysvr/scripts/CommDef"
	"gatewaysvr/scripts/ConnectionMgr"
	"gatewaysvr/scripts/config"
	"lgameserver/pkg/cs"
	"lgameserver/pkg/ss"
	"lgameserver/src/base/logger"
	"lgameserver/src/base/relay"
	"lgameserver/src/base/utils"
)

var XXTeaKey []byte

func DealMsgFromGamesvr2Client(head *ss.SSHead, msg utils.IMessage) {
	if XXTeaKey == nil {
		XXTeaKey = []byte(config.Conf.Common.XXTeaKey)
	}
	logger.Log.Debug("sshead============%v", head)
	ClientMsg := msg.(*ss.SSSend2Client)
	var cshead cs.CSHead
	err := cshead.Unmarshal(ClientMsg.GetCsHeadBuff())
	if err != nil {
		logger.Log.Error("DealMsgFromGamesvr2Client error:%s", err.Error())
		return
	}
	var userConn *CommDef.UserConnector
	//登录协议返回
	if cshead.GetCmdId() == uint32(cs.EnmCmdValue_ECV_CSLoginRes) {
		var loginRes cs.CSLoginRes
		err = loginRes.Unmarshal(ClientMsg.GetCsMsgBuff())
		if err != nil {
			logger.Log.Error("loginRes.Unmarshal error:%s", err.Error())
			return
		}
		logger.Log.Debug("loginRes==%v", loginRes)
		if cshead.GetRetCode() == uint32(cs.EnmRetCode_SUCCESS) {
			//绑定ACCID
			connId := cshead.Context
			logger.Log.Debug("CSLoginRes connId== %d", connId)
			userConn = ConnectionMgr.ConnMgr.GetUserConnByConnId(connId)
			if userConn != nil {
				ConnectionMgr.ConnMgr.BindAccId(userConn, loginRes.GetPlayerData().GetAccId())
				//发送上线协议
				SendOnline2OnlineSvr(loginRes.GetPlayerData().GetAccId())
			}
		}
	} else {
		accId := head.GetAccId()
		userConn = ConnectionMgr.ConnMgr.GetUserConnByAccid(accId)
	}
	if userConn != nil && userConn.GnetConn != nil {
		sendBuf := make([]byte, 0, len(ClientMsg.GetCsHeadBuff())+len(ClientMsg.GetCsMsgBuff())+2)
		headLenBuf := make([]byte, 2, 2)
		//登录协议不加密
		if userConn.EncryptKey != nil {
			headLenBuf[0] = 1
		} else {
			headLenBuf[0] = 0
		}

		headLenBuf[1] = byte(len(ClientMsg.GetCsHeadBuff()))
		sendBuf = append(sendBuf, headLenBuf...)
		if userConn.EncryptKey != nil {

			var encryptedBuff []byte
			if cshead.GetCmdId() == uint32(cs.EnmCmdValue_ECV_CSLoginRes) {
				LoginRes := &cs.CSLoginRes{}
				LoginRes.Unmarshal(ClientMsg.GetCsMsgBuff())
				LoginRes.EncryptKey = []byte(userConn.EncryptKey)
				buffMsg, _ := LoginRes.Marshal()
				msgBuff := append(ClientMsg.GetCsHeadBuff(), buffMsg...)
				encryptedBuff = utils.XXTeaEncryptAndCompress(msgBuff, XXTeaKey)
			} else {
				msgBuff := append(ClientMsg.GetCsHeadBuff(), ClientMsg.GetCsMsgBuff()...)
				encryptedBuff = utils.XXTeaEncryptAndCompress(msgBuff, userConn.EncryptKey)
			}

			sendBuf = append(sendBuf, encryptedBuff...)

		} else {
			sendBuf = append(sendBuf, ClientMsg.GetCsHeadBuff()...)
			sendBuf = append(sendBuf, ClientMsg.GetCsMsgBuff()...)
		}
		(*userConn.GnetConn).AsyncWrite(sendBuf)
	} else {
		logger.Log.Error("---------------------------------userConn == nil")
	}

}

func SendOnline2OnlineSvr(accId uint64) {
	ssHead := &ss.SSHead{
		CmdId:    uint32(ss.EnmCmdValue_ECV_SSOnlineReq),
		SeqId:    0,
		Src:      config.GetSelfConnIdx(),
		Dst:      uint64(ss.ServiceType_OnlineSvr),
		External: 0,
		AccId:    accId,
	}
	onlineReq := &ss.SSOnlineReq{
		AccId:        accId,
		SrcServiceId: config.GetSelfConnIdx(),
	}
	relay.TransMsg(config.GetSelfConnIdx(), uint64(ss.ServiceType_OnlineSvr), accId, ssHead, onlineReq)
}

func SendOffline2OnlineSvr(accId uint64) {
	ssHead := &ss.SSHead{
		CmdId:    uint32(ss.EnmCmdValue_ECV_SSOfflineReq),
		SeqId:    0,
		Src:      config.GetSelfConnIdx(),
		Dst:      uint64(ss.ServiceType_OnlineSvr),
		External: 0,
		AccId:    accId,
	}
	offlineReq := &ss.SSOfflineReq{
		AccId:        accId,
		SrcServiceId: config.GetSelfConnIdx(),
	}
	relay.TransMsg(config.GetSelfConnIdx(), uint64(ss.ServiceType_OnlineSvr), accId, ssHead, offlineReq)
}
*/
