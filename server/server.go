package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/fmyxyz/connectYou/server/data"
	"github.com/fmyxyz/connectYou/server/listen"
)

var (
	//28656、35580
	port = flag.Int("port", 35580, "侦听端口号。")
)

func main() {
	flag.Parse()
	listen.Handler0.AddHandlerFunc(initHandler)
	listen.Listen(*port)
}

func initHandler(conn net.Conn, data data.Metadata) data.Metadata {
	var bs []byte
	bs = make([]byte, 1)
	i, err := conn.Read(bs)
	if err != nil {
		log.Print("连接初始化失败：", err)
	}
	if i == 1 {
		fmt.Println("请求为：", bs)
		return data.Metadata{Data: bs, ReqType: bs[0]}
	}
	return data.Metadata{}
}
