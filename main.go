package main

import (
	httphandler "github.com/Sirlanri/distiot-user-server/httpHandler"
	"github.com/Sirlanri/distiot-user-server/rpcHandler/tokenrpc"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

func main() {
	log.Log.Debugln("hello user后端")

	go httphandler.IrisInit()
	go tokenrpc.RPCListen()

	//tokenrpc.RPCTest("test1 failed")
	//tokenrpc.RPCTest("d703fcc1-655e-4a4f-bdb1-5fecd89b07cb")
	//从书中学到的奇淫技巧(*^▽^*)
	select {}

}
