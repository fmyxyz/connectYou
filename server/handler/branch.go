package handler

import (
	"bufio"

	"github.com/fmyxyz/connectYou/server/data"
)

type BranchHandler struct {
	handlers map[byte]Handler
}

func NewBranchHandler() BranchHandler {
	return BranchHandler{handlers: make(map[byte]Handler, 1<<6)}
}
func (bh *BranchHandler) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, data data.Metadata) data.Metadata {
	h := bh.handlers[data.ReqType]
	data = h.Handle(bufReader, bufWriter, data)
	return data
}
func (bh *BranchHandler) AddHandler(b byte, h Handler) {
	bh.handlers[b] = h
}

func (bh *BranchHandler) AddHandlerFunc(b byte, h HandlerFunc) {
	bh.handlers[b] = h
}
