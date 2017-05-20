package handler

import "net"

type ConnectHandler struct {
}

func NewConnectHandler() ConnectHandler {
	return ConnectHandler{}
}
func (bh *ConnectHandler) Handle(conn net.Conn, data Metadata) Metadata {
	return data
}
