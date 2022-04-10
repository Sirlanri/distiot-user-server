package main

import (
	httphandler "github.com/Sirlanri/distiot-user-server/httpHandler"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

func main() {
	log.Log.Debugln("hello")

	go httphandler.IrisInit()
	//从书中学到的奇淫技巧(*^▽^*)
	select {}

}
