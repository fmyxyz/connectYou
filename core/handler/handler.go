package handler

import (
	"bufio"
	//"net"

	"github.com/fmyxyz/connectYou/core/data"
)

type Handler interface {
	Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, data *data.Metadata)
}

type HandlerFunc func(bufReader *bufio.Reader, bufWriter *bufio.Writer, data *data.Metadata)

func (handlerFunc HandlerFunc) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, data *data.Metadata) {
	handlerFunc(bufReader, bufWriter, data)
}
