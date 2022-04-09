package httphandler

import (
	"github.com/Sirlanri/distiot-user-server/server/encrypt"
	"github.com/Sirlanri/distiot-user-server/server/log"
	"github.com/Sirlanri/distiot-user-server/server/user"
	"github.com/kataras/iris/v12"
)

//handler 处理登录请求
func LoginHandler(con iris.Context) {
	//http传入的用户数据
	var userRec UserLoginData
	err := con.ReadJSON(&userRec)
	if err != nil {
		log.Log.Infoln("httpHandler LoginHandler 读取json失败", err.Error())
		con.StatusCode(401)
		return
	}
	//数据库中用户的信息
	userDB, err := user.GetUserByMailMysql(userRec.Mail)
	if err == nil || !encrypt.ComparePW(userDB.Pw, userRec.Pw) {
		con.StatusCode(401)
		log.Log.Infoln("httpHandler LoginHandler 登录失败", err.Error())
		return
	}
	//进行权限的设置
	sess := encrypt.Sess.Start(con)
	sess.Set("user_id", userDB.ID)
	sess.Set("user_level", userDB.Level)

	con.StatusCode(200)
	return

}

//handler 处理注册请求
func RegisterHandler(con iris.Context) {
	var userRec UserRegisterData
	err := con.ReadJSON(&userRec)
	if err != nil {
		log.Log.Infoln("httpHandler RegisterHandler 读取json失败", err.Error())
		con.StatusCode(401)
		return
	}

	//验证传入数据合法性
	if userRec.Mail == "" || userRec.Vcode == "" || userRec.Pw == "" {
		con.WriteString("邮箱密码验证码均不可为空")
		con.StatusCode(401)
		return
	}
	err = user.UserRegisterVerify(userRec.Mail, userRec.Vcode)
	if err != nil {
		con.WriteString(err.Error())
		con.StatusCode(401)
		return
	}
	err = user.InsertUserMysql(userRec.Mail, userRec.Pw)
	if err != nil {
		con.WriteString("服务器错误，注册失败")
		con.StatusCode(401)
		return
	}
}

//handler 申请发送验证码
func ApplyVcodeHandler(con iris.Context) {
	mail := con.URLParam("mail")
	if mail == "" {
		con.StatusCode(401)
		log.Log.Infoln("httpHandler ApplyVcodeHandler 参数错误")
		return
	}
	vcode, err := user.VcodeProcess(mail)
	if err != nil {
		con.StatusCode(401)
		log.Log.Infoln("httpHandler ApplyVcodeHandler 生成验证码失败", err.Error())
		return
	}
	err = user.SendMailVcode(vcode, mail)
	if err != nil {
		con.StatusCode(401)
		return
	}
	con.StatusCode(200)
}

//handler 处理重置密码请求
func ResetPW(con iris.Context) {

}

//用户登录数据
type UserLoginData struct {
	Mail string `json:"mail"`
	Pw   string `json:"pw"`
}

//用户注册时传入的数据结构
type UserRegisterData struct {
	Mail  string `json:"mail"`
	Pw    string `json:"pw"`
	Vcode string `json:"vcode"`
}
