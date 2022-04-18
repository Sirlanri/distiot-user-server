package main

import (
	"github.com/Sirlanri/distiot-user-server/rpcHandler/rpcTest"
	"github.com/Sirlanri/distiot-user-server/rpcHandler/rpcserver"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

func main() {
	log.Log.Debugln("hello user后端")

	//go httphandler.IrisInit()
	go rpcserver.RPCListen()

	//tokenrpc.RPCTest("test1 failed")
	rpcTest.RPCTest("d703fcc1-655e-4a4f-bdb1-5fecd89b07cb")
	rpcTest.RPCTest2(6)
	rpcTest.RPCTest2(5)
	rpcTest.RPCTest2(4)
	//从书中学到的奇淫技巧(*^▽^*)
	select {}

}
