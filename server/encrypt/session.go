package encrypt

import (
	"time"

	"github.com/Sirlanri/distiot-user-server/server/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/sessions/sessiondb/redis"
)

var Sess *sessions.Sessions

func init() {
	SessionInit()
}

func SessionInit() {
	db := redis.New(redis.Config{
		Network:   "tcp",
		Addr:      config.Config.RedisAddr,
		Timeout:   2 * time.Second,
		MaxActive: 10,
		Database:  "1",
		Prefix:    "iris-",
		//Delim:     "iris-",
		//Driver:    GoRedis(), // redis.Radix() can be used instead.
	})
	iris.RegisterOnInterrupt(func() {
		db.Close()
	})

	//全局Session
	Sess = sessions.New(sessions.Config{
		Cookie: "sessionscookieid",
		//AllowReclaim:    true,
		CookieSecureTLS: true,
	})
	Sess.UseDatabase(db)
}
