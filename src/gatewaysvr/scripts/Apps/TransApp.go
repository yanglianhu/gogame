package Apps

/*import (
	"lgameserver/src/base/logger"

	cs "lgameserver/pkg/cs"
	ss "lgameserver/pkg/ss"
)

func HandleTransReq(head *ss.SSHead, buffer []byte) {
	subBitMarkOffset := 4
	subHeadLenOffset := subBitMarkOffset + 1
	subHeadOffset := subHeadLenOffset + 1
	subHeadLen := int(buffer[subHeadLenOffset])
	handler := GetTransmitHandle(head.SubCmdId)
	logger.Log.Debug("HandleTransReq----------------------------%s", head.String())
	if handler == nil {
		logger.Log.Error("not found trans cmd handler. Head: %s", head.String())
		return
	}

	if head.SubCmdId < uint32(ss.EnmCmdValue_ECV_SSRegistReq) {
		// 客户端消息转发
		subHead := &cs.CSHead{}
		subMsg := cs.NewClientMessage(int32(head.SubCmdId))

		if subMsg == nil {
			logger.Log.Error("unpack trans not found message. CmdId: %d", head.SubCmdId)
			return
		}

		if err := subHead.Unmarshal(buffer[subHeadOffset : subHeadOffset+subHeadLen]); err != nil {
			logger.Log.Error("unpack trans subHead failed. Err: %v, CmdId: %d", err, head.SubCmdId)
			return
		}

		subMsgOffset := subHeadOffset + subHeadLen
		if err := subMsg.Unmarshal(buffer[subMsgOffset:]); err != nil {
			logger.Log.Error("unpack trans subMsg failed. Err: %v, CmdId: %d", err, head.SubCmdId)
			return
		}

		logger.Log.Debug("recv transmit message. Head: %s, SubHead: %s, SubMsg: %s", head.String(), subHead.String(), subMsg.String())

		handler(head, subHead, subMsg)
	} else {
		// 服务器消息转发
		subHead := &ss.SSHead{}
		subMsg := ss.NewServiceMessage(int32(head.SubCmdId))

		if subMsg == nil {
			logger.Log.Error("unpack trans not found message. CmdId: %d", head.SubCmdId)
			return
		}

		if err := subHead.Unmarshal(buffer[subHeadOffset : subHeadOffset+subHeadLen]); err != nil {
			logger.Log.Error("unpack trans subHead failed. Err: %v, CmdId: %d", err, head.SubCmdId)
			return
		}

		subMsgOffset := subHeadOffset + subHeadLen
		if err := subMsg.Unmarshal(buffer[subMsgOffset:]); err != nil {
			logger.Log.Error("unpack trans subMsg failed. Err: %v, CmdId: %d", err, head.SubCmdId)
			return
		}

		logger.Log.Debug("recv transmit message. Head: %s, SubHead: %s, SubMsg: %s", head.String(), subHead.String(), subMsg.String())

		handler(head, subHead, subMsg)
	}
}*/
