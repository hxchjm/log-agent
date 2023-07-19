package server

import (
	"log"
	"net"
	"os"
)

const (
	_address = "/var/run/lancer/collector.sock"
)

func Init() {
	// 创建Unix域数据报套接字
	addr := &net.UnixAddr{Name: _address, Net: "unixgram"}
	conn, err := net.ListenUnixgram("unixgram", addr)
	if err != nil {
		log.Println("创建Unix域数据报套接字失败：", err)
		os.Exit(1)
	}
	defer conn.Close()

	go process(conn)
}

func process(conn *net.UnixConn) {
	for {
		// 接收消息
		buf := make([]byte, 1024)
		n, ua, err := conn.ReadFromUnix(buf)
		if err != nil {
			log.Println("接收消息失败：", err)
			os.Exit(1)
		}
		log.Printf("接收到消息：(%+v),ua:(%+v)", string(buf[:n]), ua)
	}
}
