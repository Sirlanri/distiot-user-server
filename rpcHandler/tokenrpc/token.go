package tokenrpc

import (
	"context"
	"errors"
	"net"

	pb "github.com/Sirlanri/distiot-user-server/rpcHandler/pb"
	"github.com/Sirlanri/distiot-user-server/server/config"
	"github.com/Sirlanri/distiot-user-server/server/device"
	"github.com/Sirlanri/distiot-user-server/server/log"
	"github.com/Sirlanri/distiot-user-server/server/token"
	"google.golang.org/grpc"
)

type server struct {
	pb.TokenGuideServer
}

//token RPC服务
func (s *server) GetToken(ctx context.Context, req *pb.Tokenmsg) (*pb.UserIDmsg, error) {
	log.Log.Debugln("RPC服务-GetToken 传入", req.String())
	id, err := token.GetUserIDByToken(req.Token)
	if err != nil {
		log.Log.Warnln("RPC服务 获取userID失败", err.Error())
		return nil, errors.New("Token错误")
	}
	return &pb.UserIDmsg{UserID: int64(id)}, nil
}

func (s *server) GetdIDByToken(ctx context.Context, req *pb.Tokenmsg) (*pb.DeviceIDsmsg, error) {
	log.Log.Debugln("RPC服务-GetdIDByToken 传入", req.String())
	id, err := device.GetAllDeviceIDByToken(req.Token)
	if err != nil {
		log.Log.Warnln("RPC服务 获取设备ID失败", err.Error())
		return nil, errors.New("Token错误")
	}
	ids := make([]int64, 0, len(*id))
	for _, v := range *id {
		ids = append(ids, int64(v))
	}
	return &pb.DeviceIDsmsg{DeviceIDs: ids}, nil
}

//RPC的监听服务 阻塞式的，应该用协程启动
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
