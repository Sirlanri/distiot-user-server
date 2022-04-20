package httphandler

import (
	"strconv"

	"github.com/Sirlanri/distiot-user-server/server/device"
	"github.com/kataras/iris/v12"
)

//handler 创建新设备的handler。header中需要包含token，Parma包含设备名称
func CreateDeviceHandler(con iris.Context) {
	token := con.GetHeader("token")
	dName := con.URLParam("dname")
	dataType := con.URLParamIntDefault("dataType", 0)
	if token == "" || dName == "" || dataType == 0 {
		con.StatusCode(401)
		con.WriteString("传入数据不合法，token和设备名不能为空")
		return
	}
	did, err := device.CreateDevice(token, dName, dataType)
	if err != nil {
		con.StatusCode(500)
		con.WriteString("服务器错误，创建失败")
		return
	}
	con.StatusCode(200)
	con.WriteString(strconv.Itoa(did))
	return
}
