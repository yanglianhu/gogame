package Apps

/*import (
	ss "lgameserver/pkg/ss"
	"lgameserver/src/base/logger"

	"github.com/walkon/gnet"
)

type HandleService func(conn gnet.Conn, head *ss.SSHead, iMsg interface{})
type HandleMisc func(conn gnet.Conn, head *ss.SSHead, miscReq *ss.SSMiscReq, iMsg interface{})

type HandleTransmit func(head *ss.SSHead, subHead interface{}, subMsg interface{})

var (
	handleServiceMap  map[ss.EnmCmdValue]HandleService
	handleTransmitMap map[uint32]HandleTransmit
	handleMiscMap     map[string]HandleMisc
)

func init() {
	handleServiceMap = make(map[ss.EnmCmdValue]HandleService)
	handleTransmitMap = make(map[uint32]HandleTransmit)
	handleMiscMap = make(map[string]HandleMisc)
}

func RegistServiceHandle(cmdId ss.EnmCmdValue, handle HandleService) int {
	if _, ok := handleServiceMap[cmdId]; ok {
		logger.Log.Error("duplicate register cmd. CmdId: %d", cmdId)
		return -1
	}
	logger.Log.Debug("register cmd success. CmdId: %d", cmdId)
	handleServiceMap[cmdId] = handle
	return 0
}

func GetServiceHandle(cmdId uint32) HandleService {
	if v, ok := handleServiceMap[ss.EnmCmdValue(cmdId)]; !ok {
		return nil
	} else {
		return v
	}
}

func RegistMiscHandle(msgName string, handle HandleMisc) int {
	if _, ok := handleMiscMap[msgName]; ok {
		logger.Log.Error("duplicate register misc cmd. MsgName: %s", msgName)
		return -1
	}
	logger.Log.Debug("register misc cmd success. MsgName: %s", msgName)
	handleMiscMap[msgName] = handle
	return 0
}

func GetMiscHandle(msgName string) HandleMisc {
	if v, ok := handleMiscMap[msgName]; !ok {
		return nil
	} else {
		return v
	}
}

func RegistTransmitHandle(cmdId uint32, handle HandleTransmit) int {
	if _, ok := handleTransmitMap[cmdId]; ok {
		logger.Log.Error("duplicate register transmit cmd. CmdId: %d", cmdId)
		return -1
	}
	logger.Log.Debug("register transmit cmd success. CmdId: %d", cmdId)
	handleTransmitMap[cmdId] = handle
	return 0
}

func GetTransmitHandle(cmdId uint32) HandleTransmit {
	if v, ok := handleTransmitMap[cmdId]; !ok {
		return nil
	} else {
		return v
	}
}*/
