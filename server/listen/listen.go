package listen

import (
	"log"
	"net"

	"github.com/fmyxyz/connectYou/server/handler"
)

func Listen(port int) {
	laddr := &net.TCPAddr{Port: port}
	listenner, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Println("侦听端口错误：", err)
		return
	}
	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Println("获取连接错误：", err)
			return
		}
		go accept(conn)
	}
}

var Handler0 handler.BaseHandler = handler.NewBaseHandler()

func accept(conn net.Conn) {
	Handler0.Handle(conn, handler.Metadata{})
}
