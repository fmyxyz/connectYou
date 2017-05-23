package handler

import (
	"net"

	"github.com/fmyxyz/connectYou/server/data"
)

type BranchHandler struct {
	Handlers map[byte]Handler
}

func NewBranchHandler() BranchHandler {
	return BranchHandler{Handlers: make(map[byte]Handler, 1<<6)}
}
func (bh *BranchHandler) Handle(conn net.Conn, data data.Metadata) data.Metadata {
	h := bh.Handlers[data.ReqType]
	data = h.Handle(conn, data)
	return data
}
func (bh *BranchHandler) AddHandler(b byte, h Handler) {
	bh.Handlers[b] = h
}

func (bh *BranchHandler) AddHandlerFunc(b byte, h HandlerFunc) {
	bh.Handlers[b] = h
}
