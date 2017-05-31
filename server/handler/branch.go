package handler

import (
	"bufio"

	"github.com/fmyxyz/connectYou/server/data"
)

type BranchHandler struct {
	handlers map[uint8]Handler
	SequenceHandler
}

func NewBranchHandler() BranchHandler {
	return BranchHandler{handlers: make(map[uint8]Handler, 1<<6), SequenceHandler: NewSequenceHandler()}
}
func (bh *BranchHandler) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, data *data.Metadata) {
	h := bh.handlers[data.ReqType]
	if h != nil {
		h.Handle(bufReader, bufWriter, data)
	} else {
		bh.SequenceHandler.Handle(bufReader, bufWriter, data)
	}
}
func (bh *BranchHandler) AddHandler(b uint8, h Handler) {
	bh.handlers[b] = h
}

func (bh *BranchHandler) AddHandlerFunc(b uint8, h HandlerFunc) {
	bh.handlers[b] = h
}
