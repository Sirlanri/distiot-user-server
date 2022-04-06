package db

import (
	"github.com/Sirlanri/distiot-user-server/server/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	//MySQL数据库指针
	Mdb *gorm.DB
)

func init() {
	connectMysqlByGorm()
	connectRedis()
}

func connectMysqlByGorm() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/distiot-user?parseTime=true"
	Mdb, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Log.Errorln("server-db MySQL连接失败", err.Error())
		return
	}
	err = Mdb.Error
	if err != nil {
		log.Log.Errorln("server-db MySQL ping失败", err.Error())
		return
	}
	log.Log.Infoln("server-db MySQL连接成功")

}
