package device

import (
	"github.com/Sirlanri/distiot-user-server/server/db"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

//将deviceID-userID-deviceName存入mysql，返回该设备ID
func InsertUserDeviceMysql(userID int, name string) (int, error) {
	var device Device
	device.Uid = userID
	device.Dname = name
	err := db.Mdb.Create(&device).Error
	if err != nil {
		log.Log.Warnln("server-device InsertUserDeviceMysql 存入MySQL失败", err.Error())
		return 0, err
	}
	return device.Did, nil
}

//将单个用户Token-deviceID集合存入Redis
func InsertUserDeviceRedis(userToken string, deviceID int) error {
	_, err := db.Rdb.SAdd(db.RedisCtx, userToken, deviceID).Result()
	if err != nil {
		log.Log.Warnln("server-device InsertUserDeviceRedis 存入Redis失败", err.Error())
		return err
	}
	return nil
}

//数据库中的device模型
type Device struct {
	Did   int    `gorm:"primary_key"`
	Uid   int    `gorm:"int"`
	Dname string `gorm:"varchar(255)"`
}
