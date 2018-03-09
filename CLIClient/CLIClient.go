package main

import (
	"bufio"
	"context"
	"encoding/json"
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
	conn, err := net.Dial("tcp", "127.0.0.1:35580")
	if err != nil {
		log.Println("连接不上主机：", err)
		return
	}
	defer conn.Close()

	i := strings.LastIndex(conn.LocalAddr().String(), ":")
	from_user_id = conn.LocalAddr().String()[i+1:]
	log.Println("from_user_id", from_user_id)
	key := fmt.Sprint(rand.Int63n(time.Now().UnixNano()))
	context.WithValue(ctx, core.CONNKEY, key)
	context.WithValue(ctx, core.CONN, conn)
	context.WithValue(ctx, core.USERID, from_user_id)
	go request.SendHeartbeatHandler(ctx)

	login(ctx)

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

func login(ctx context.Context) {
	key := core.GetConnKey(ctx)
	conn := core.GetConn(ctx)
	var bigDataLength int32 = 1 << 17
	for {
		md := data.NewMetadata(bigDataLength, key)
		md.ReqType = listen.JsonHandlerKey
		m := data.Message{MsgType: "login", Msg: "登录", FromConnId: key}
		md.Data, _ = json.Marshal(m)
		md.Length = int32(len(md.Data))
		dl, err := conn.Write(md.Packing())
		if err != nil || dl != 6 {
			log.Println("发送数据错误：", err)
			break
		}
		time.Sleep(10 * time.Minute)
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
		chatWith(argss, bufReader, md, key, conn)
	case exit:
		break
	}
}

func chatWith(argss string, bufReader *bufio.Reader, md *data.Metadata, key string, conn net.Conn) {
	//聊天模式
	fmt.Println("和" + argss + "聊天中。。。输入send发送，bye结束聊天。")
	var msg string
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
				return
			}
		default:
			msg += msgLine
		}
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
