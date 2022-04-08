package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Sirlanri/distiot-user-server/server/config"
	"github.com/Sirlanri/distiot-user-server/server/encrypt"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

/* 这个文件包括用户注册登录重置密码等请求
换言之，就是未进入系统时的操作*/

//传入用户邮箱，生成验证码，并写入Redis
func VcodeProcess(mail string) (string, error) {
	code := encrypt.GenerateVcode()
	err := InsertVcodeRedis(mail, code)
	if err != nil {
		return "", err
	}
	return code, nil
}

//发送验证码邮件
func SendMailVcode(vode, mail string) error {
	var bodyData VcodeMail
	str := "您的验证码是： " + vode + " ，请在5分钟内使用，如非本人操作，请忽略此邮件。"
	bodyData.Content = str
	bodyData.Addr = mail
	bodyData.Topic = "Distiot验证码"
	body, err := json.Marshal(bodyData)
	if err != nil {
		log.Log.Warnln("server-user SendMailVcode json.Marshal 失败", err.Error())
		return err
	}
	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequest("POST", "https://api.ri-co.cn/mail/send", bytes.NewReader(body))
	if err != nil {
		log.Log.Warnln("server-user SendMailVcode http.NewRequest 失败", err.Error())
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	//想不到吧，email服务器我设置了认证
	req.Header.Set("token", config.Config.EmailToken)
	res, err := client.Do(req)
	if err != nil {
		log.Log.Warnln("server-user SendMailVcode client.Do 失败", err.Error())
		return err
	}
	if res.StatusCode != 200 {
		log.Log.Warnln("server-user SendMailVcode 远程mail服务器响应失败", res.StatusCode)
		return errors.New("server-user SendMailVcode 远程mail服务器响应失败")
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

//发送验证码所需的结构体
type VcodeMail struct {
	Addr    string `json:"addr"`
	Topic   string `json:"topic"`
	Content string `json:"content"`
}
