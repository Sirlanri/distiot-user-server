package tokenrpc

import (
	"context"
	"time"

	"github.com/Sirlanri/distiot-user-server/rpcHandler/pb"
	"github.com/Sirlanri/distiot-user-server/server/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RPCTest() {
	log.Log.Debugln("RPC测试开始")
	conn, err := grpc.Dial("localhost:8092",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Log.Errorln("测试失败！无法连接端口", err.Error())
		return
	}
	defer conn.Close()
	c := pb.NewTokenGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetToken(ctx, &pb.Tokenmsg{
		Token: "This is Token",
	})
	log.Log.Debugln("返回消息：", r.UserID)

}
