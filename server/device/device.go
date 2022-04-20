package device

import (
	"github.com/Sirlanri/distiot-user-server/server/db"
	"github.com/Sirlanri/distiot-user-server/server/token"
)

/* 用户创建新设备
传入token，调用GetUserIDByToken获取userID，将映射关系存入redis的集合，
返回设备的ID
*/
func CreateDevice(userToken, dName string, userid, dataType int) (int, error) {
	//插入数据库
	did, err := InsertUserDeviceMysql(userid, int(dataType), dName)
	if err != nil {
		return 0, err
	}
	//插入redis
	InsertUserDeviceRedis(userToken, did)

	return did, nil

}

//传入token，获取该用户的设备列表
func GetAllDeviceIDByToken(userToken string) (*[]int, error) {
	//判断用户Token是否真实存在
	userID, err := token.GetUserIDByToken(userToken)
	if err != nil {
		return nil, err
	}
	//从redis中获取
	didList, err := GetAllDeviceIDByTokenRedis(userToken)
	if err != nil {
		//去mysql中搜索
		didList, err = GetAllDeviceIDByUserIDMysql(userID)
		if err != nil {
			return nil, err
		}
		//将结果存入redis
		InsertUserTokenDeviceRedis(userToken, *didList...)
	}
	return didList, nil
}

//传入用户ID，获取该用户全部的device信息
func GetAllDevicesByIDMysql(userID int) (*[]Device, error) {
	var devices []Device
	db.Mdb.Where("uid=?", userID).Find(&devices)
	return &devices, nil
}
