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
		log.Log.Errorln("failed to listen: %v", err)
	}

	s1 := grpc.NewServer()
	pb.RegisterDeviceServiceServer(s1, &deviceServer{})
	pb.RegisterTokenGuideServer(s1, &tokenServer{})
	log.Log.Debugln("RPC server1 listening at %v", lis.Addr())
	if err := s1.Serve(lis); err != nil {
		log.Log.Errorln("failed to serve s1: %v", err)
	}

	log.Log.Debugln("RPC server2 listening at %v", lis.Addr())
	if err := s1.Serve(lis); err != nil {
		log.Log.Errorln("failed to serve s2: %v", err)
	}
}
