package device

import (
	"errors"
	"strconv"

	"github.com/Sirlanri/distiot-user-server/server/db"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

//将deviceID-userID-deviceName存入mysql，返回该设备ID
func InsertUserDeviceMysql(userID, dataType int, name string) (int, error) {
	var device Device
	device.Uid = userID
	device.Dname = name
	device.Datatype = int32(dataType)
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

//从Redis中获取token对应的deviceID集合
func GetAllDeviceIDByTokenRedis(userToken string) (*[]int, error) {
	deviceIDList := make([]int, 0)
	ids := db.Rdb.SMembers(db.RedisCtx, userToken).Val()
	for index, id := range ids {
		log.Log.Debugln(index, id)
		deviceID, err := strconv.Atoi(id)
		if err != nil {
			log.Log.Warnln("server-device GetAllDeviceIDByTokenRedis 获取Redis中的Token-dID集合 转换失败", err.Error())
			return nil, err
		}
		deviceIDList = append(deviceIDList, deviceID)
	}
	if len(deviceIDList) == 0 {
		log.Log.Warnln("server-device GetAllDeviceIDByTokenRedis 从Redis中获取设备ID列表失败")
		return nil, errors.New("找不到集合")
	}
	return &deviceIDList, nil
}

//传入用户id，获取用户的所有设备ID
func GetAllDeviceIDByUserIDMysql(userID int) (*[]int, error) {
	deviceIDList := make([]int, 0)
	err := db.Mdb.Table("devices").Select("did").Where("uid = ?", userID).Find(&deviceIDList).Error
	if err != nil {
		log.Log.Warnln("server-device GetAllDeviceIDByUserIDMysql 从MySQL中获取设备ID列表失败", err.Error())
		return nil, err
	}
	return &deviceIDList, nil
}

//将用户Token-deviceID集合存入Redis，删除旧的集合
func InsertUserTokenDeviceRedis(userToken string, deviceID ...int) error {
	err := db.Rdb.Del(db.RedisCtx, userToken).Err()
	if err != nil {
		log.Log.Warnln("server-device InsertUserTokenDeviceRedis 删除Redis中的Token-dID集合失败", err.Error())
		return err
	}
	//可变参数传递
	for _, id := range deviceID {
		_, err = db.Rdb.SAdd(db.RedisCtx, userToken, id).Result()
		if err != nil {
			log.Log.Warnln("server-device InsertUserTokenDeviceRedis 存入Redis中的Token-dID集合失败", err.Error())
			return err
		}
	}

	return nil
}

//从MySQL中获取设备的数据类型
func GetTypeMysql(deviceID int) (int32, error) {
	var device Device
	device.Did = deviceID
	err := db.Mdb.First(&device).Error
	if err != nil {
		log.Log.Warnln("server-device GetTypeMysql 从MySQL中获取设备数据类型失败", err.Error())
		return 0, err
	}
	return device.Datatype, nil
}

//数据库中的device模型
type Device struct {
	Did      int    `gorm:"primary_key"`
	Uid      int    `gorm:"int"`
	Dname    string `gorm:"varchar(255)"`
	Datatype int32  `gorm:"tinyint"`
}
