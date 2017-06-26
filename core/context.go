package core

import (
	"context"
	"net"
)

const (
	CONN    = "net.Conn"
	CONNKEY = "ConnKey"
	USERID  = "UserId"
)

func GetConn(ctx context.Context) net.Conn {
	i := ctx.Value(CONN)
	if i != nil {
		conn, ok := i.(net.Conn)
		if ok {
			return conn
		} else {
			return nil
		}
	}
	return nil
}
func GetConnKey(ctx context.Context) string {
	i := ctx.Value(CONNKEY)
	if i != nil {
		conn, ok := i.(string)
		if ok {
			return conn
		} else {
			return ""
		}
	}
	return ""
}
func GetUserId(ctx context.Context) string {
	i := ctx.Value(USERID)
	if i != nil {
		conn, ok := i.(string)
		if ok {
			return conn
		} else {
			return ""
		}
	}
	return ""
}

type ConnContext struct {
	Conn    net.Conn
	ConnKey string
	UserId  string
	Ctx     context.Context
	Cancel  context.CancelFunc
}
