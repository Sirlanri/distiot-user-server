package main

import (
	"github.com/Sirlanri/distiot-user-server/server/device"
	"github.com/Sirlanri/distiot-user-server/server/log"
)

func main() {
	log.Log.Debugln("hello")

	//go httphandler.IrisInit()
	ids, err := device.GetAllDeviceIDByToken("703fcc1-655e-4a4f-bdb1-5fecd89b07cb")
	if err != nil {
		log.Log.Errorln(err)
	}
	log.Log.Debugln(ids)
	//从书中学到的奇淫技巧(*^▽^*)
	select {}

}
