package device

import "github.com/Sirlanri/distiot-user-server/server/token"

/* 用户创建新设备
传入token，调用GetUserIDByToken获取userID，将映射关系存入redis的集合，
返回设备的ID
*/
func CreateDevice(userToken, dName string) (int, error) {
	userid, err := token.GetUserIDByToken(userToken)
	if err != nil {
		return 0, err
	}
	//插入数据库
	did, err := InsertUserDeviceMysql(userid, dName)
	if err != nil {
		return 0, err
	}
	//插入redis
	InsertUserDeviceRedis(userToken, did)

	return did, nil

}
