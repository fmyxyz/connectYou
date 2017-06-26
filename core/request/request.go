package request

import (
	"context"
	"log"
	"time"

	"github.com/fmyxyz/connectYou/core"
	"github.com/fmyxyz/connectYou/core/data"
	"github.com/fmyxyz/connectYou/server/listen"
)

//心跳
//
//		十分钟一次
func SendHeartbeatHandler(ctx context.Context) {
	key := core.GetConnKey(ctx)
	conn := core.GetConn(ctx)
	var bigDataLength int32 = 1 << 17
	for {
		md := data.NewMetadata(bigDataLength, key)
		md.ReqType = listen.HeartbeatHandlerKey
		md.Length = int32(len(md.Data))
		dl, err := conn.Write(md.Packing())
		if err != nil || dl != 6 {
			log.Println("发送数据错误：", err)
			break
		}
		time.Sleep(10 * time.Minute)
	}
}
