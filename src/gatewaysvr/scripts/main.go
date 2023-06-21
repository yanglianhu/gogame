package main

import (
	"gogame/protobuf/cs"
	"log"
	"net"
)

func main() {
	port := "8080" // 网关服务器监听的端口

	// 监听指定端口
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Error starting the server:", err)
	}

	log.Println("Server started. Listening on port", port)

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// 处理客户端连接
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var head cs.CSHead
	// 处理连接逻辑
	log.Println("New client connected:", conn.RemoteAddr())
	// 读取和处理客户端消息
	buffer := make([]byte, 1024)
	for {
		// 读取客户端消息
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading from client:", err)
			break
		}

		// 处理客户端消息
		msg := string(buffer[:n])
		log.Println("Received message from client:", msg)

		// 响应客户端消息
		response := []byte("Message received: " + msg)
		_, err = conn.Write(response)
		if err != nil {
			log.Println("Error sending response to client:", err)
			break
		}
	}

	log.Println("Client disconnected:", conn.RemoteAddr())
}
