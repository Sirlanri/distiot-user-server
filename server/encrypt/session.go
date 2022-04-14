package encrypt

import (
	"time"

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
		Addr:      "127.0.0.1:6379",
		Timeout:   2 * time.Second,
		MaxActive: 10,
		Password:  "",
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
		Cookie:       "sessionscookieid",
		AllowReclaim: true,
	})
	Sess.UseDatabase(db)
}
