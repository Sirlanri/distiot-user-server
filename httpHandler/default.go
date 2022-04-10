package httphandler

import (
	"github.com/Sirlanri/distiot-user-server/server/config"
	"github.com/kataras/iris/v12"
)

func IrisInit() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	user := app.Party("/user").AllowMethods(iris.MethodOptions)
	{
		user.Post("/login", LoginHandler)
		user.Post("/register", RegisterHandler)
		user.Get("/applyvcode", ApplyVcodeHandler)
		user.Post("/resetpw", ResetPWHandler)
		user.Get("/createdevice", CreateDeviceHandler)
	}
	var portStr = config.Config.HttpPort
	app.Run(iris.Addr(":" + portStr))
}
