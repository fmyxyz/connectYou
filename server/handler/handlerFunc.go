package handler

import "net"

type HandlerFunc func(conn net.Conn, data Metadata) Metadata

func (handlerFunc HandlerFunc) Handle(conn net.Conn, data Metadata) Metadata {
	return handlerFunc(conn, data)
}
