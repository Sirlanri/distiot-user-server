package encrypt

import (
	"github.com/Sirlanri/distiot-user-server/server/log"
	"github.com/google/uuid"
)

func Getuuid() {
	id := uuid.New()
	log.Log.Debugln("uuid:", id)
}
