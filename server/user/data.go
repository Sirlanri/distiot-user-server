package user

import (
	"time"

	"github.com/Sirlanri/distiot-user-server/server/db"
	"github.com/Sirlanri/distiot-user-server/server/encrypt"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

//包含id、mail、pw
type User struct {
	ID    string `gorm:"primary_key"`
	Mail  string `gorm:"varChar(255)"`
	Pw    string `gorm:"varChar(255)"`
	Token string `gorm:"varChar(255)"`
	Dnum  int    `gorm:"default:10"`
	Level int    `gorm:"int"`
}

//传入userID，获取用户信息
func GetUserInfoByIDMysql(userID int) (*User, error) {
	var user User
	err := db.Mdb.Where("id = ?", userID).First(&user).Error
	if err != nil {
		log.Log.Warnln("server-user GetUserByIDMysql 从MySQL获取用户信息失败", err.Error())
		return nil, err
	}
	return &user, nil
}

//传入邮箱 从MySQL中获取用户的信息
func GetUserByMailMysql(mail string) (*User, error) {
	user := new(User)
	err := db.Mdb.Where("mail = ?", mail).First(&user).Error
	if err != nil {
		log.Log.Infoln("server-user GetPwByMailMysql 查询失败", err.Error())
		return nil, err
	}
	return user, nil
}

//检查该邮箱是否存在，存在就返回true，不存在就返回false
func MailExist(mail string) bool {
	var user User
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
		log.Log.Warnln("server-user InsertVcode 存入Redis失败", err.Error())
		return err
	}
	return nil
}

//从Redis中获取验证码
func GetVcodeRedis(mail string) (string, error) {
	var code string
	code, err := db.Rdb.Get(db.RedisCtx, mail).Result()
	if err != nil {
		log.Log.Warnln("server-user GetVcodeRedis 从Redis获取验证码失败", err.Error())
		return "", err
	}
	return code, nil
}

//传入mail，pw。自动生成token，将用户信息插入数据库
func InsertUserMysql(mail, pw string) error {
	var user User
	user.Mail = mail
	hashed, err := encrypt.GenerateHash(pw)
	if err != nil {
		log.Log.Warnln("server-user InsertUserMysql 加密失败", err.Error())
		return err
	}
	user.Pw = hashed
	user.Token = encrypt.Generateuuid() //uuid无法生成时，会返回""
	user.Level = 1
	err = db.Mdb.Omit("ID").Create(&user).Error
	if err != nil {
		log.Log.Warnln("server-user InsertUserMysql 存入MySQL失败", err.Error())
		return err
	}
	return nil
}

//更新用户密码。传入用户邮箱，未加密的密码
func UpdatePW(usermail string, pw string) error {
	var user User
	hashed, err := encrypt.GenerateHash(pw)
	if err != nil {
		log.Log.Warnln("server-user UpdatePW 加密失败", err.Error())
		return err
	}
	user.Pw = hashed
	err = db.Mdb.Table("users").Where("mail = ?", usermail).Update("pw", hashed).Error
	if err != nil {
		log.Log.Warnln("server-user UpdatePW 更新MySQL失败", err.Error())
		return err
	}
	return nil
}
