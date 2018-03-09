package route

import (
	"bufio"

	"github.com/fmyxyz/connectYou/core/data"
)

type Router interface {
	Router(bufReader *bufio.Reader, bufWriter *bufio.Writer, msg *data.Message)
}

//纯json
type RouteHandler struct {
	RouterMap map[string]Router
}

func NewRouteHandler() *RouteHandler {
	hh := RouteHandler{}
	return &hh
}
func (bh *RouteHandler) Router(bufReader *bufio.Reader, bufWriter *bufio.Writer, msg *data.Message) {
	router := bh.RouterMap[msg.MsgType]
	if router != nil {
		router.Router(bufReader, bufWriter, msg)
	}
}
