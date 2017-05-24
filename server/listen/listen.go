package listen

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/fmyxyz/connectYou/server/data"
	"golang.org/x/net/websocket"
)

func StartTCPServer(port int) {
	add := &net.TCPAddr{Port: port}
	lt, err := net.ListenTCP("tcp", add)
	if err != nil {
		log.Println("侦听端口错误：", err)
		return
	}
	for {
		conn, err := lt.AcceptTCP()
		conn.SetKeepAlive(true)
		conn.SetKeepAlivePeriod(30 * time.Second)
		if err != nil {
			log.Println("获取连接错误：", err)
			return
		}
		go handleConn(conn)
	}
}

func StartWebsocketServer(port int) {
	http.Handle("/ws", websocket.Handler(websocketHandler))
	err := http.ListenAndServe(":"+string(port), nil)
	if err != nil {
		log.Println("侦听端口错误：", err)
		return
	}
}

func websocketHandler(ws *websocket.Conn) {
	handleConn(ws)
}

func handleConn(conn net.Conn) {
	//TODO 保存于客户端的链接
	defer conn.Close()
	var bigDataLength int32 = 1 << 17
	for {
		bufReader := bufio.NewReader(conn)
		bufWriter := bufio.NewWriter(conn)
		md := data.NewMetadata(bigDataLength)
		err := md.Unpacking(bufReader)
		if err != nil {
			log.Println(err)
			conn.Close()
			break
		}
		md.Conn = conn
		routeHandler.Handle(bufReader, bufWriter, md)
	}
}
