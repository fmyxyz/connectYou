package handler

import (
	"net"

	"github.com/fmyxyz/connectYou/server/data"
)

type HandlerFunc func(conn net.Conn, data data.Metadata) data.Metadata

func (handlerFunc HandlerFunc) Handle(conn net.Conn, data data.Metadata) data.Metadata {
	return handlerFunc(conn, data)
}
