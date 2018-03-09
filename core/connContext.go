package core

import (
	"context"
	"net"
)

type ConnContext struct {
	Session
	context.Context
}

func NewConnContext(conn net.Conn) *ConnContext {
	conncontext := &ConnContext{Context: GetAppContext(), Session: NewSession(conn)}
	GetAppContext().conns[conncontext.id] = conncontext
	return conncontext
}
