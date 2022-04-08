package user

import (
	"time"

	"github.com/Sirlanri/distiot-user-server/server/db"
	"github.com/Sirlanri/distiot-user-server/server/encrypt"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

//包含id、mail、pw
type UserData struct {
	ID    string `gorm:"primary_key"`
	Mail  string `gorm:"varChar(255)"`
	Pw    string `gorm:"varChar(255)"`
	Token string `gorm:"varChar(255)"`
	Level int    `gorm:"int"`
}

//传入邮箱 从MySQL中获取用户的信息
func GetUserByMailMysql(mail string) (*UserData, error) {
	user := new(UserData)
	err := db.Mdb.Where("mail = ?", mail).First(&user).Error
	if err != nil {
		log.Log.Infoln("server-db GetPwByMailMysql 查询失败", err.Error())
		return nil, err
	}
	return user, nil
}

//检查该邮箱是否存在，存在就返回true，不存在就返回false
func MailExist(mail string) bool {
	var user UserData
	err := db.Mdb.Where("mail = ?", mail).First(&user).Error
	if err != nil {
		return false
	}
	return true
}

//将验证码存入Redis
func InsertVcodeRedis(mail, code string) error {
	_, err := db.Rdb.Set(db.RedisCtx, mail, code, time.Minute*5).Result()
	if err != nil {
		log.Log.Warnln("server-db InsertVcode 存入Redis失败", err.Error())
		return err
	}
	return nil
}

//从Redis中获取验证码
func GetVcodeRedis(mail string) (string, error) {
	var code string
	code, err := db.Rdb.Get(db.RedisCtx, mail).Result()
	if err != nil {
		log.Log.Warnln("server-db GetVcodeRedis 从Redis获取验证码失败", err.Error())
		return "", err
	}
	return code, nil
}

//将用户信息插入数据库
func InsertUserMysql(mail, pw string) error {
	var user UserData
	user.Mail = mail
	user.Pw = pw
	user.Token = encrypt.Generateuuid() //uuid无法生成时，会返回""
	err := db.Mdb.Create(&user).Error
	if err != nil {
		log.Log.Warnln("server-db InsertUserMysql 存入MySQL失败", err.Error())
		return err
	}
	return nil
}
