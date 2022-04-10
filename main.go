package main

import (
	"time"

	"github.com/Sirlanri/distiot-user-server/rpcHandler/tokenrpc"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

func main() {
	log.Log.Debugln("hello user后端")

	//go httphandler.IrisInit()
	go func() {
		tokenrpc.RPCListen()
	}()
	time.Sleep(time.Second * 3)

	tokenrpc.RPCTest()
	//从书中学到的奇淫技巧(*^▽^*)
	select {}

}
