package CommDef

import "net"

const (
	ConnectionStat_Authed int = 1
)

type UserConnector struct {
	TcpConn    *net.Conn
	Status     int
	Uid        uint64
	CreateTime int64
	Context    int64
	LastSeqId  uint32
	EncryptKey []byte
}
