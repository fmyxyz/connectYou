package core

import (
	"context"
	"errors"
	"time"
)

type ConnContext struct {
	Session
	context.Context
}

func NewConnContext(parent context.Context) *ConnContext {
	return &ConnContext{Context: parent}
}

func (cc *ConnContext) Deadline() (deadline time.Time, ok bool) {
	return cc.Context.Deadline()
}

func (cc *ConnContext) Done() <-chan struct{} {
	return cc.Context.Done()
}

func (cc *ConnContext) Err() error {
	return errors.New("context error")
}

func (cc *ConnContext) Value(key interface{}) interface{} {
	if k, ok := key.(string); ok {
		return cc.GetAttr(k)
	}
	return nil
}
