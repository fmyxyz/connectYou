package main

import (
	"flag"

	"github.com/fmyxyz/connectYou/server/listen"
)

var (
	//28656、35580
	port = flag.Int("port", 35580, "侦听端口号。")
)

func main() {
	flag.Parse()
	listen.StartTCPServer(35580)
	listen.StartWebsocketServer(28656)
}
