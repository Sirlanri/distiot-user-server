package httphandler

import (
	"github.com/Sirlanri/distiot-user-server/server/config"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func IrisInit() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	app.Use(crs)
	app.OnErrorCode(iris.StatusNotFound, NotFound)
	user := app.Party("/user").AllowMethods(iris.MethodOptions, iris.MethodGet, iris.MethodPost)
	{
		user.Post("/login", LoginHandler)
		user.Get("/logout", LogoutHandler)
		user.Post("/register", RegisterHandler)
		user.Get("/applyvcode", ApplyVcodeHandler)
		user.Post("/resetpw", ResetPWHandler)
		user.Get("/createdevice", CreateDeviceHandler)
		user.Get("/getdevices", GetDevicesHandler)
		user.Get("/getdeviceids", GetDeviceIDsHandler)
		user.Get("/getuserinfo", GetUserInfoHandler)
	}
	var portStr = config.Config.HttpPort
	app.Run(iris.Addr(":" + portStr))
}
