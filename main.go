package main

import (
	httphandler "github.com/Sirlanri/distiot-user-server/httpHandler"
	"github.com/Sirlanri/distiot-user-server/rpcHandler/rpcserver"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

func main() {
	log.Log.Debugln("hello user后端")

	//go httphandler.IrisInit()
	go rpcserver.RPCListen()

	go httphandler.IrisInit()

	//从书中学到的奇淫技巧(*^▽^*)
	select {}

}
