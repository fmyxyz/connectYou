package handler

import (
	"bufio"
	"log"
	"net"

	"github.com/fmyxyz/connectYou/server/data"
)

type HeartbeatHandler struct {
	conn net.Conn
}

func NewHeartbeatHandler(second int) HeartbeatHandler {
	hh := HeartbeatHandler{}
	return hh
}
func (bh *HeartbeatHandler) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, md data.Metadata) data.Metadata {
	bh.conn = md.Conn
	_, err := bufWriter.Write(md.Packing())
	if err != nil {
		log.Println(err)
		md.Conn.Close()
	}
	return md
}
