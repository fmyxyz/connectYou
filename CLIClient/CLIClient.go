package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/fmyxyz/connectYou/core"
	"github.com/fmyxyz/connectYou/core/data"
	"github.com/fmyxyz/connectYou/core/request"
	"github.com/fmyxyz/connectYou/server/listen"
)

var (
	port         = flag.Int("port", 35580, "连接端口")
	host         = flag.String("host", "127.0.0.1", "连接主机地址")
	server       = flag.String("server", "127.0.0.1:35580", "连接主机地址和端口")
	from_user_id string
	ctx          = context.Background()
)

const (
	//连接到服务器
	connnect = "conn"
	//添加朋友
	add = "add"
	//查看朋友列表
	list = "list"
	//删除朋友
	del = "del"
	//和朋友聊天
	chat = "chat"
	//发送信息
	send = "send"
	//取消
	cancel = "cnacel"
	//退出聊天
	bye = "bye"
	//退出系统
	exit = "exit"
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *server)
	i := strings.LastIndex(conn.LocalAddr().String(), ":")
	from_user_id = conn.LocalAddr().String()[i+1:]
	log.Println("from_user_id", from_user_id)
	if err != nil {
		log.Println("连接不上主机：", err)
		return
	}
	defer conn.Close()
	key := fmt.Sprint(rand.Int63n(time.Now().UnixNano()))
	context.WithValue(ctx, core.CONNKEY, key)
	context.WithValue(ctx, core.CONN, conn)
	context.WithValue(ctx, core.USERID, from_user_id)
	go request.SendHeartbeatHandler(ctx)
	bufReader := bufio.NewReader(os.Stdin)
	for {
		//读取命令行中一行数据
		cmd, err := bufReader.ReadString(byte('\n'))
		cmd = strings.TrimSpace(cmd)
		cmds := strings.Split(cmd, ":")
		cmd = cmds[0]
		var argss string
		if len(cmds) > 1 {
			argss = cmds[1]
		}
		if err == nil {
			startChat(cmd, argss, bufReader, key, conn)
		} else {
			break
		}
	}
}

//开始聊天
func startChat(cmd string, argss string, bufReader *bufio.Reader, key string, conn net.Conn) {
	var bigDataLength int32 = 1 << 17
	md := data.NewMetadata(bigDataLength, key)
	switch cmd {
	case connnect:
	case add:
	case list:
	case del:
	case chat:
		//聊天模式
		fmt.Println("和" + argss + "聊天中。。。输入send发送，bye结束聊天。")
		var msg string
	CHAT:
		for {
			msgLine, _ := bufReader.ReadString(byte('\n'))
			switch strings.TrimSpace(msgLine) {
			case send:
				sendChatMsg(md, msg, key, argss, conn)
				fmt.Println("信息已发送。。。输入bye结束聊天。")
				msg = ""
			case cancel:
				msg = ""
			case bye:
				if len(msg) == 0 {
					fmt.Println("与" + argss + "的聊天结束。。。")
					break CHAT
				}
			default:
				msg += msgLine
			}
		}
	case exit:
		break
	}
}

//发送聊天数据
func sendChatMsg(md *data.Metadata, msg string, key string, argss string, conn net.Conn) {
	md.ReqType = listen.JsonHandlerKey
	m := data.Message{MsgType: "chat", Msg: msg, FromConnId: key, ToUserId: argss}
	data, _ := json.Marshal(m)
	md.Data = data
	md.Length = int32(len(md.Data))
	bs := md.Packing()
	dl, err := conn.Write(bs)
	if err != nil || dl != int(md.Length)+6 {
		log.Println("发送数据错误：", err)
	}
}
