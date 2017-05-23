package handler

import (
	"github.com/fmyxyz/connectYou/server/data"
	"net"
)

type ConnectHandler struct {
}

func NewConnectHandler() ConnectHandler {
	return ConnectHandler{}
}
func (bh *ConnectHandler) Handle(conn net.Conn, data data.Metadata) data.Metadata {
	return data
}
