package main

import (
	//httphandler "github.com/Sirlanri/distiot-user-server/httpHandler"
	"github.com/Sirlanri/distiot-user-server/server/log"
	"github.com/Sirlanri/distiot-user-server/server/token"
)

func main() {
	log.Log.Debugln("hello")
	id, _ := token.GetUserIDByToken("d703fc1-655e-4a4f-bdb1-5fecd89b07cb")
	log.Log.Debugln("id:", id)
	id, _ = token.GetUserIDByToken("d703fcc1-655e-4a4f-bdb1-5fecd89b07cb")
	log.Log.Debugln("id:", id)
	id, _ = token.GetUserIDByToken("d703fcc1-655e-4a4f-bdb1-5fecd89b07cb")
	log.Log.Debugln("id:", id)
	//httphandler.IrisInit()
	//从书中学到的奇淫技巧(*^▽^*)
	select {}

}
