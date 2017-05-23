package listen

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/fmyxyz/connectYou/server/data"
)

func StartServer(port int) {
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

func handleConn(conn net.Conn) {
	defer conn.Close()
	bs := make([]byte, 0, 1<<17)
	buff := bytes.NewBuffer(bs)
	md, err := data.Unpacking(buff)
	if err != nil {
		if err.(data.DataTooLengthError) {
			//conn.Close()
		} else if err.(data.ProtocolError) {
			conn.Close()
		} else {
			conn.Close()
		}
	} else {
		initHandler(conn, md)
	}
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
