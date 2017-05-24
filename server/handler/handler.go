package handler

import (
	"bufio"

	"github.com/fmyxyz/connectYou/server/data"
)

type Handler interface {
	Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, data data.Metadata) data.Metadata
}

type HandlerFunc func(bufReader *bufio.Reader, bufWriter *bufio.Writer, data data.Metadata) data.Metadata

func (handlerFunc HandlerFunc) Handle(bufReader *bufio.Reader, bufWriter *bufio.Writer, data data.Metadata) data.Metadata {
	return handlerFunc(bufReader, bufWriter, data)
}
