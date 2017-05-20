package main

import (
	"flag"
	"log"
	"net"
)

var (
	port   = flag.Int("port", 234, "连接端口")
	host   = flag.String("host", "127.0.0.1", "连接主机地址")
	server = flag.String("server", "127.0.0.1:7070", "连接主机地址和端口")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *server)
	if err != nil {
		log.Println("连接不上主机：", err)

	}
	conn.Write([]byte{1})
}
