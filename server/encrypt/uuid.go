package encrypt

import (
	"github.com/google/uuid"
)

func Generateuuid() string {
	id := uuid.New()
	return id.String()
}
