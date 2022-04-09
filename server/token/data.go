package token

import (
	"errors"
	"strconv"

	"github.com/Sirlanri/distiot-user-server/server/db"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

//传入token，从redis中读取userID
func GetUserIDByTokenRedis(token string) (userID int, err error) {
	userIDStr, err := db.Rdb.HGet(db.RedisCtx, "userid", token).Result()
	if err != nil {
		log.Log.Warnln("server-token GetUserIDByToken 从Redis获取用户ID失败", err.Error())
		return 0, err
	}
	userID, err = strconv.Atoi(userIDStr)
	if err != nil {
		log.Log.Warnln("server-token GetUserIDByToken 转换用户ID失败", err.Error())
		return 0, err
	}
	return userID, nil
}

//传入token，从MySQL中读取userID
func GetUserIDByTokenMysql(token string) (userID int, err error) {
	//err = db.Mdb.First(&userID, "token = ?", token).Error
	db.Mdb.Table("users").Select("id").Where("token = ?", token).Scan(&userID)
	if userID == 0 {
		log.Log.Warnln("server-token GetUserIDByToken 从MySQL获取用户ID失败")
		return 0, errors.New("无此Token")
	}
	return userID, nil
}

//将Token-UserID映射关系存入Redis
func InsertTokenUserIDRedis(token string, userid int) error {
	_, err := db.Rdb.HSet(db.RedisCtx, "userid", token, strconv.Itoa(userid)).Result()
	if err != nil {
		log.Log.Warnln("server-token InsertTokenUserID 存入Redis失败", err.Error())
		return err
	}
	return nil
}
