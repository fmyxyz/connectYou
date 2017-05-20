package handler

import "net"

type SequenceHandler struct {
	Handlers []Handler
}

func NewSequenceHandler() SequenceHandler {
	return SequenceHandler{Handlers: make([]Handler, 0, 1<<6)}
}
func (bh *SequenceHandler) Handle(conn net.Conn, data Metadata) Metadata {
	for _, h := range bh.Handlers {
		data = h.Handle(conn, data)
	}
	return data
}
func (bh *SequenceHandler) AddHandler(h Handler) {
	bh.Handlers = append(bh.Handlers, h)
}

func (bh *SequenceHandler) AddHandlerFunc(h HandlerFunc) {
	bh.Handlers = append(bh.Handlers, h)
}
