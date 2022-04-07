package httphandler

import (
	"github.com/Sirlanri/distiot-user-server/server/log"
	"github.com/Sirlanri/distiot-user-server/server/user"
	"github.com/kataras/iris/v12"
)

//handler 处理登录请求
func LoginHandler(con iris.Context) {
	var userData UserLoginData
	err := con.ReadJSON(&userData)
	if err != nil {
		log.Log.Infoln("httpHandler LoginHandler 读取json失败", err.Error())
		con.StatusCode(401)
		return
	}
	if user.UserLogin(userData.Mail, userData.Pw) {

		con.StatusCode(200)
	} else {
		con.StatusCode(401)
		log.Log.Infoln("httpHandler LoginHandler 登录失败", err.Error())
		return
	}

}

//handler 处理注册请求
func RegisterHandler(con iris.Context) {

}

//handler 处理重置密码请求
func ResetPW(con iris.Context) {

}

//用户登录数据
type UserLoginData struct {
	Mail string `json:"mail"`
	Pw   string `json:"pw"`
}
