package Apps

/*import (
	"gatewaysvr/scripts/ConnectionMgr"
	ss "lgameserver/pkg/ss"
	"lgameserver/src/base/logger"
	relay "lgameserver/src/base/relay"
)

func init() {
	RegistServiceHandle(ss.EnmCmdValue_ECV_SSRegistRes, relay.HandleServRegistRes)
	RegistServiceHandle(ss.EnmCmdValue_ECV_SSHeartRes, relay.HandleServHeartRes)
	RegistTransmitHandle(uint32(ss.EnmCmdValue_ECV_SSPlayerLogout), HandlePlayerLogout)
	RegistTransmitHandle(uint32(ss.EnmCmdValue_ECV_SSBanUser), HandleBanUser)
}
func HandlePlayerLogout(head *ss.SSHead, subHead interface{}, subMsg interface{}) {
	req := subMsg.(*ss.SSPlayerLogout)
	accId := req.AccId
	userConn := ConnectionMgr.ConnMgr.GetUserConnByAccid(accId)
	if userConn != nil && userConn.GnetConn != nil {
		logger.Log.Debug("close connection by kick------------------------------")
		(*userConn.GnetConn).Close()
	}

}

func HandleBanUser(head *ss.SSHead, subHead interface{}, subMsg interface{}) {
	req := subMsg.(*ss.SSBanUser)
	accId := req.AccId
	userConn := ConnectionMgr.ConnMgr.GetUserConnByAccid(accId)
	if userConn != nil && userConn.GnetConn != nil {
		logger.Log.Debug("close connection by kick------------------------------")
		(*userConn.GnetConn).Close()
	}
}*/
