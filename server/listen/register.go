package listen

import "github.com/fmyxyz/connectYou/server/handler"

var routeHandler handler.BranchHandler

func init() {
	routeHandler = handler.NewBranchHandler()
	routeHandler.AddHandler(byte(1), handler.NewHeartbeatHandler(30))
}
