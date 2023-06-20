package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

//将整数转换成大端字节序的数组
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//发送消息
func sendMsg(tcpconn net.Conn, data string) {
	rawMsgLen := len(data)
	rawBytes := []byte(data)
	sendBuff := append(IntToBytes(rawMsgLen), rawBytes...)
	tcpconn.Write(sendBuff)
}

//接收消息，阻塞模式
func recvMsg(tcpconn net.Conn) (error, []byte) {
	reader := bufio.NewReader(tcpconn)
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.BigEndian, &length)
	if err != nil {
		return err, nil
	}
	pack := make([]byte, int(4+length))
	_, err = io.ReadFull(reader, pack)
	if err != nil {
		return err, nil
	}
	return nil, pack[4:]
}

func main() {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:9003", 5*time.Second)
	if err != nil {
		fmt.Println("DialTimeout error:" + err.Error())
		return
	}
	sendMsg(conn, "i have a dream")
	err, msg := recvMsg(conn)
	if err == nil {
		fmt.Println("recv msg " + string(msg))
	}

}
