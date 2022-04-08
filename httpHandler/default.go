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
		user.Get("/vcode", ApplyVcodeHandler)
		user.Post("/resetpw", ResetPW)
	}
	var portStr = config.Config.Port
	app.Run(iris.Addr(":" + portStr))
}
