package rpcserver

import (
	"net"

	pb "github.com/Sirlanri/distiot-user-server/rpcHandler/pb"
	"github.com/Sirlanri/distiot-user-server/server/config"
	"github.com/Sirlanri/distiot-user-server/server/log"
	"google.golang.org/grpc"
)

//RPC的监听服务 阻塞式的，应该用协程启动
func RPCListen() {
	addr := "localhost:" + config.Config.RpcPort
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Log.Errorln("RPC 无法监听端口", err)
	}

	s := grpc.NewServer()
	pb.RegisterAllRpcServerServer(s, &servers{})

	log.Log.Debugln("RPC server 监听端口：", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Log.Errorln("RPC服务启动失败", err)
	}

}
