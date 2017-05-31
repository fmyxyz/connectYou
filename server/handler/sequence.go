package handler

import (
	"bufio"

	"github.com/fmyxyz/connectYou/server/data"
)

type SequenceHandler struct {
	Handlers []Handler
}

func NewSequenceHandler() SequenceHandler {
	return SequenceHandler{Handlers: make([]Handler, 0, 1<<6)}
}
func (bh *SequenceHandler) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, data *data.Metadata) {
	for _, h := range bh.Handlers {
		h.Handle(bufReader, bufWriter, data)
	}

}
func (bh *SequenceHandler) AddHandler(h Handler) {
	bh.Handlers = append(bh.Handlers, h)
}

func (bh *SequenceHandler) AddHandlerFunc(h HandlerFunc) {
	bh.Handlers = append(bh.Handlers, h)
}
