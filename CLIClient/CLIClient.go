package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/fmyxyz/connectYou/server/data"
	"github.com/fmyxyz/connectYou/server/listen"
)

var (
	port   = flag.Int("port", 35580, "连接端口")
	host   = flag.String("host", "127.0.0.1", "连接主机地址")
	server = flag.String("server", "127.0.0.1:35580", "连接主机地址和端口")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *server)
	if err != nil {
		log.Println("连接不上主机：", err)
		return
	}
	key := fmt.Sprint(rand.Int63n(time.Now().UnixNano()))
	var bigDataLength int32 = 1 << 17
	for {
		log.Println("纯json数据打包。。。")
		md := data.NewMetadata(bigDataLength, key)
		//md.ReqType = listen.HeartbeatHandlerKey
		//conn.Write(md.Packing())
		md.ReqType = listen.JsonHandlerKey
		md.Data = []byte(`{
	"msg_type": "msg_type",
	"msg": "msg",
	"from_user_id": "from_user_id",
	"to_user_id": "to_user_id",
	"from_conn_id": "from_conn_id",
	"to_conn_id": "to_conn_id"
}`)
		md.Length = int32(len(md.Data))
		log.Println("纯json数据开始发送。。。")
		dl, err := conn.Write(md.Packing())
		if err != nil {
			conn.Close()
			log.Println("发送数据错误：", err)
			break
		}
		log.Println("纯json数据已发送", dl, "bytes")
		time.Sleep(10 * time.Second)
	}

}
