package httphandler

import "github.com/kataras/iris/v12"

func IrisInit() {
	app := iris.New()
	user := app.Party("/user").AllowMethods(iris.MethodOptions)
	{
		user.Post("/login", LoginHandler)
		user.Post("/register", RegisterHandler)
		user.Post("/resetpw", ResetPW)
	}
}