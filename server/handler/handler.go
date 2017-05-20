package handler

import "net"

type Handler interface {
	Handle(conn net.Conn, data Metadata) Metadata
}

type BaseHandler struct {
	Handlers []Handler
}

func NewBaseHandler() BaseHandler {
	return BaseHandler{Handlers: make([]Handler, 0, 1<<6)}
}
func (bh *BaseHandler) Handle(conn net.Conn, data Metadata) Metadata {
	for _, h := range bh.Handlers {
		data = h.Handle(conn, data)
	}
	return data
}
func (bh *BaseHandler) AddHandler(h Handler) {
	bh.Handlers = append(bh.Handlers, h)
}

func (bh *BaseHandler) AddHandlerFunc(h HandlerFunc) {
	bh.Handlers = append(bh.Handlers, h)
}
