package rpcserver

import (
	"context"
	"errors"

	pb "github.com/Sirlanri/distiot-user-server/rpcHandler/pb"
	"github.com/Sirlanri/distiot-user-server/server/device"
	"github.com/Sirlanri/distiot-user-server/server/log"
	"github.com/Sirlanri/distiot-user-server/server/token"
)

type tokenServer struct {
	pb.TokenGuideServer
}

//token RPC服务
func (s *tokenServer) GetUidByToken(ctx context.Context, req *pb.Tokenmsg) (*pb.UserIDmsg, error) {
	log.Log.Debugln("RPC服务-GetUidByToken 传入", req.String())
	id, err := token.GetUserIDByToken(req.Token)
	if err != nil {
		log.Log.Warnln("RPC服务 获取userID失败", err.Error())
		return nil, errors.New("Token错误")
	}
	return &pb.UserIDmsg{UserID: int64(id)}, nil
}

func (s *tokenServer) GetdIDByToken(ctx context.Context, req *pb.Tokenmsg) (*pb.DeviceIDsmsg, error) {
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
