package listen

import (
	"net"

	"github.com/fmyxyz/connectYou/server/handler"
)

var routeHandler handler.BranchHandler

const (
	//心跳
	HeartbeatHandlerKey = uint8(1)
	//纯文本
	StringHandlerKey = uint8(2)
	//纯Json
	JsonHandlerKey = uint8(3)
	//富文本 json
	RTFHandlerKey = uint8(4)
	//二进制
	BinaryHandlerKey = uint8(5)
)

func init() {
	routeHandler = handler.NewBranchHandler()
	//心跳
	routeHandler.AddHandler(HeartbeatHandlerKey, handler.NewHeartbeatHandler(30))
	routeHandler.AddHandler(JsonHandlerKey, handler.NewJsonHandler())
	routeHandler.AddHandler(RTFHandlerKey, handler.NewRTFHandler())
	handler.ConnYouMap = make(map[string]net.Conn, 10000)
}
