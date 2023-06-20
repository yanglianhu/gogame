package main

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/panjf2000/gnet"
)

type Server struct {
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
	codec        gnet.ICodec
}

func (this *Server) OnInitComplete(svr gnet.Server) (action gnet.Action) {
	fmt.Println("OnInitComplete")
	return
}

func (this *Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println("OnOpened")
	return
}

func (this *Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	fmt.Println("OnClosed")
	return
}

func (this *Server) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	fmt.Println("React-----------")
	c.AsyncWrite(frame)
	return
}

func (this *Server) Tick() (delay time.Duration, action gnet.Action) {
	fmt.Println("Tick-----------")
	return time.Second, gnet.Action(0)
}

func StartServe() {
	reuseport := true
	multicore := false
	async := true
	network := "tcp"
	addr := "127.0.0.1:9003"

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
	ms := &Server{
		network:   network,
		addr:      addr,
		multicore: multicore,
		async:     async,
		codec:     codec,
	}

	fmt.Println("start server service ...")

	gnet.Serve(ms, network+"://"+addr,
		gnet.WithLockOSThread(async),
		gnet.WithMulticore(multicore),
		gnet.WithReusePort(reuseport),
		gnet.WithTicker(true),
		gnet.WithTCPKeepAlive(time.Minute*1),
		gnet.WithCodec(codec),
		gnet.WithTCPNoDelay(gnet.TCPDelay),
		gnet.WithLoadBalancing(gnet.SourceAddrHash),
		// gnet.WithLogger(logger.Log),
	)
}

func main() {
	StartServe()
}
