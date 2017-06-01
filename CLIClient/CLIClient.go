package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	"strings"

	"github.com/fmyxyz/connectYou/core/data"
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
	defer conn.Close()
	key := fmt.Sprint(rand.Int63n(time.Now().UnixNano()))
	go sendHeartbeatHandler(key, conn)
	bufReader := bufio.NewReader(os.Stdin)
	for {
		cmd, err := bufReader.ReadString(byte('\n'))
		cmd = strings.TrimSpace(cmd)
		if err == nil {
			if cmd == "exit" {
				break
			}
			var bigDataLength int32 = 1 << 17
			md := data.NewMetadata(bigDataLength, key)
			md.ReqType = listen.JsonHandlerKey
			m := data.Message{MsgType: "test", Msg: cmd, FromConnId: key}
			data, _ := json.Marshal(m)
			md.Data = data
			md.Length = int32(len(md.Data))
			bs := md.Packing()
			dl, err := conn.Write(bs)
			if err != nil || dl != int(md.Length)+6 {
				log.Println("发送数据错误：", err)
				break
			}
		} else {
			break
		}
	}
}

func sendHeartbeatHandler(key string, conn net.Conn) {
	var bigDataLength int32 = 1 << 17
	for {
		md := data.NewMetadata(bigDataLength, key)
		md.ReqType = listen.HeartbeatHandlerKey
		md.Length = int32(len(md.Data))
		dl, err := conn.Write(md.Packing())
		if err != nil || dl != 6 {
			log.Println("发送数据错误：", err)
			break
		}
		time.Sleep(10 * time.Second)
	}
}
