package tokenrpc

import (
	"context"
	"net"

	pb "github.com/Sirlanri/distiot-user-server/rpcHandler/pb"
	"github.com/Sirlanri/distiot-user-server/server/config"
	"github.com/Sirlanri/distiot-user-server/server/log"
	"google.golang.org/grpc"
)

type server struct {
	pb.TokenGuideServer
}

//token RPC服务
func (s *server) GetToken(ctx context.Context, req *pb.Tokenmsg) (*pb.UserIDmsg, error) {
	log.Log.Debugln("RPC服务 传入", req.String())
	return &pb.UserIDmsg{UserID: 2}, nil
}

//RPC的监听服务 阻塞式的
func RPCListen() {
	addr := "localhost:" + config.Config.RpcPort
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Log.Errorln("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTokenGuideServer(s, &server{})
	log.Log.Debugln("RPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Log.Errorln("failed to serve: %v", err)
	}
}
