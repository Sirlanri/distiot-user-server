package rpcserver

import (
	"context"
	"errors"

	pb "github.com/Sirlanri/distiot-user-server/rpcHandler/pb"
	"github.com/Sirlanri/distiot-user-server/server/device"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

type deviceServer struct {
	pb.DeviceServiceServer
}

//device数据类型 RPC服务
func (s *deviceServer) GetDataTypeBydID(ctx context.Context, req *pb.Didmsg) (*pb.DataTypemsg, error) {
	log.Log.Debugln("RPC服务-GetDataTypeBydID 传入", req.String())
	id, err := device.GetTypeMysql(int(req.Id))
	if err != nil {
		log.Log.Warnln("RPC服务 获取数据类型失败", err.Error())
		return nil, errors.New("设备ID错误")
	}
	return &pb.DataTypemsg{DataType: id}, nil
}
