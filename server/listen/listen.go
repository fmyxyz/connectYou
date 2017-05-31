package listen

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/fmyxyz/connectYou/server/data"
	"github.com/fmyxyz/connectYou/server/handler"
	"golang.org/x/net/websocket"
)

func StartTCPServer(port int) {
	add := &net.TCPAddr{Port: port}
	lt, err := net.ListenTCP("tcp", add)
	if err != nil {
		log.Println("侦听端口错误：", err)
		return
	}
	log.Println("侦听端口.....", add.Port)
	for {
		conn, err := lt.AcceptTCP()
		log.Println("接受链接：" + conn.RemoteAddr().String())
		//conn.SetKeepAlive(true)
		//conn.SetKeepAlivePeriod(30 * time.Second)
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
	//保存于客户端的链接
	key := fmt.Sprint(rand.Int63n(time.Now().UnixNano()))
	handler.ConnYouMap[key] = conn

	defer conn.Close()
	var bigDataLength int32 = 1 << 17
	bufReader := bufio.NewReader(conn)
	bufWriter := bufio.NewWriter(conn)
	for {
		log.Println("请求包处理......")
		md := data.NewMetadata(bigDataLength, key)
		log.Println("1......")
		err := md.Unpacking(bufReader)
		log.Println("2......")
		if err != nil {
			log.Println(err)
			conn.Close()
			break
		}
		md.Conn = conn
		log.Println("3......")
		routeHandler.Handle(bufReader, bufWriter, md)
		log.Println("请求包处理完成......")
	}
}
