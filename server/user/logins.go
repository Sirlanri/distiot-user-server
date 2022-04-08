package user

import (
	"errors"

	"github.com/Sirlanri/distiot-user-server/server/encrypt"
)

/* 这个文件包括用户注册登录重置密码等请求
换言之，就是未进入系统时的操作*/

//传入用户邮箱，生成验证码，并写入Redis
func VcodeProcess(mail string) error {
	code := encrypt.GenerateVcode()
	err := InsertVcodeRedis(mail, code)
	if err != nil {
		return err
	}
	return nil
}

//用户注册 传入邮箱和验证码，完成邮箱和验证码的校验
func UserRegisterVerify(mail, code string) error {
	//首先在mysql中查找，确认邮箱不存在
	if MailExist(mail) {
		return errors.New("邮箱已存在")
	}
	//核验验证码
	codeRds, err := GetVcodeRedis(mail)
	if code != codeRds || err != nil {
		return errors.New("验证码错误")
	}
	return nil
}

func UserSetSession() {

}
