package rpcTest

import (
	"context"
	"time"

	"github.com/Sirlanri/distiot-user-server/rpcHandler/pb"
	"github.com/Sirlanri/distiot-user-server/server/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RPCTest(usertoken string) {
	log.Log.Debugln("RPC测试开始")
	conn, err := grpc.Dial("localhost:8092",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Log.Errorln("RPC测试失败！无法连接端口", err.Error())
		return
	}
	defer conn.Close()
	c := pb.NewTokenGuideClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetUidByToken(ctx, &pb.Tokenmsg{
		Token: usertoken,
	})
	if err != nil {
		log.Log.Errorln("测试失败！无此Token", err.Error())
		return
	}
	log.Log.Debugln("返回消息：", r.UserID)

}

func RPCTest2(id int) {
	log.Log.Debugln("RPC测试开始")
	conn, err := grpc.Dial("localhost:8092",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Log.Errorln("RPC测试失败！无法连接端口", err.Error())
		return
	}
	defer conn.Close()
	log.Log.Debugln("开始device测试")
	c2 := pb.NewDeviceServiceClient(conn)
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()
	r2, err2 := c2.GetDataTypeBydID(ctx2, &pb.Didmsg{
		Id: int64(id),
	})
	if err2 != nil {
		log.Log.Errorln("测试失败！无此deviceID", err2.Error())
		return
	}
	log.Log.Debugln("返回消息：", r2.DataType)
}
