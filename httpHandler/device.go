package httphandler

import (
	"strconv"

	"github.com/Sirlanri/distiot-user-server/server/device"
	"github.com/Sirlanri/distiot-user-server/server/encrypt"
	"github.com/Sirlanri/distiot-user-server/server/log"
	"github.com/kataras/iris/v12"
)

//handler 创建新设备的handler。header中需要包含token，Parma包含设备名称
func CreateDeviceHandler(con iris.Context) {
	dName := con.URLParam("dname")
	dataType := con.URLParamIntDefault("dataType", 0)
	if dName == "" || dataType == 0 {
		con.StatusCode(401)
		con.WriteString("传入数据不合法，token和设备名不能为空")
		return
	}
	if dataType != 3 && dataType != 1 && dataType != 2 {
		con.StatusCode(401)
		con.WriteString("传入数据不合法，dataType只能为0,1,2")
		return
	}

	sess := encrypt.Sess.Start(con)
	userID, err := sess.GetInt("user_id")
	if err != nil {
		con.StatusCode(403)
		con.WriteString("未登录")
		return
	}
	userToken := sess.GetString("user_token")

	did, err := device.CreateDevice(userToken, dName, userID, dataType)
	if err != nil {
		con.StatusCode(500)
		con.WriteString("服务器错误，创建失败")
		return
	}
	con.StatusCode(200)
	con.WriteString(strconv.Itoa(did))
	return
}

func GetDeviceIDsHandler(con iris.Context) {
	sess := encrypt.Sess.Start(con)
	userToken := sess.GetString("user_token")
	deviceIDs, err := device.GetAllDeviceIDByToken(userToken)
	if err != nil {
		con.StatusCode(500)
		con.WriteString("服务器错误，获取失败")
		return
	}
	_, err = con.JSON(deviceIDs)
	if err != nil {
		log.Log.Warnln("httpHandler-GetDevicesHandler JSON打包失败", err)
		con.StatusCode(500)
		con.WriteString("服务器错误，获取失败")
		return
	}
}

func GetDevicesHandler(con iris.Context) {
	sess := encrypt.Sess.Start(con)
	userID, err := sess.GetInt("user_id")
	if err != nil {
		con.StatusCode(403)
		con.WriteString("未登录")
		return
	}
	devices, err := device.GetAllDevicesByIDMysql(userID)
	if err != nil {
		con.StatusCode(500)
		con.WriteString("服务器错误，获取失败")
		return
	}
	_, err = con.JSON(devices)
	if err != nil {
		log.Log.Warnln("httpHandler-GetDevicesHandler JSON打包失败", err)
		con.StatusCode(500)
		con.WriteString("服务器错误，获取失败")
		return
	}

}
