package handler

import (
	"bufio"
	"encoding/json"
	"log"
	"net"

	"github.com/fmyxyz/connectYou/core"
	"github.com/fmyxyz/connectYou/core/data"
)

type HeartbeatHandler struct {
	conn net.Conn
}

//心跳处理
func NewHeartbeatHandler(second int) *HeartbeatHandler {
	hh := HeartbeatHandler{}
	return &hh
}
func (bh *HeartbeatHandler) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, md *data.Metadata) {
	bh.conn = md.Conn
	_, err := bufWriter.Write(md.Packing())
	if err != nil {
		log.Println(err)
		md.Conn.Close()
	}
}

//纯json
type JsonHandler struct {
}

func NewJsonHandler() *JsonHandler {
	hh := JsonHandler{}
	return &hh
}
func (bh *JsonHandler) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, md *data.Metadata) {
	msg := data.Message{}
	if md.HasHandleData {
		json.Unmarshal(md.Data, &msg)
	} else {
		bs := make([]byte, md.Length)
		lens, err := bufReader.Read(bs)
		if err != nil || lens != int(md.Length) {
			md.Conn.Close()
			delete(core.ConnYouMap, md.ConnId)
			return
		}
		json.Unmarshal(md.Data, &msg)
	}

}

//富文本json
type RTFHandler struct {
}

func NewRTFHandler() *RTFHandler {
	hh := RTFHandler{}
	return &hh
}
func (bh *RTFHandler) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, md *data.Metadata) {
	_, err := bufWriter.Write(md.Packing())
	if err != nil {
		log.Println(err)
		md.Conn.Close()
	}

}

//消息分发
type RelayMsgHandler struct {
	conn net.Conn
}

func NewRelayMsgHandler(second int) RelayMsgHandler {
	hh := RelayMsgHandler{}
	return hh
}
func (bh *RelayMsgHandler) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, md *data.Metadata) {
	bh.conn = md.Conn
	_, err := bufWriter.Write(md.Packing())
	if err != nil {
		log.Println(err)
		md.Conn.Close()
	}
}
