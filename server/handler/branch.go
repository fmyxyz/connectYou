package handler

import "net"

type BranchHandler struct {
	Handlers map[byte]Handler
}

func NewBranchHandler() BranchHandler {
	return BranchHandler{Handlers: make(map[byte]Handler, 1<<6)}
}
func (bh *BranchHandler) Handle(conn net.Conn, data Metadata) Metadata {
	h := bh.Handlers[data.OperTpye]
	data = h.Handle(conn, data)
	return data
}
func (bh *BranchHandler) AddHandler(b byte, h Handler) {
	bh.Handlers[b] = h
}

func (bh *BranchHandler) AddHandlerFunc(b byte, h HandlerFunc) {
	bh.Handlers[b] = h
}
