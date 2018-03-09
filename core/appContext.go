package core

import (
	"context"
	"errors"
	"net"
	"time"
)

var appContext *AppContext

func init() {
	appContext = NewAppContext()
}
func GetAppContext() *AppContext {
	return appContext
}

type AppContext struct {
	conns map[string]*ConnContext
	context.Context
}

func NewAppContext() *AppContext {
	return &AppContext{Context: context.Background()}
}

func (cc *AppContext) Deadline() (deadline time.Time, ok bool) {
	return cc.Context.Deadline()
}

func (cc *AppContext) Done() <-chan struct{} {
	return cc.Context.Done()
}

func (cc *AppContext) Err() error {
	return errors.New("context error")
}

func (cc *AppContext) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		return cc.conns[k]
	}
	return nil
}

func (cc *AppContext) GetConn(k string) net.Conn {
	return cc.conns[k].GetConn()
}

func (cc *AppContext) RemoveConn(k string) {
	delete(cc.conns, k)
}
