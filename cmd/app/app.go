package app

import (
	"time"

	log "github.com/vessel/pkg/logger"
)

func Run() *int {
	for {
		log.Info("Starting main loop")
		a := new(int)
		time.Sleep(time.Duration(1) * time.Second)
		return a
	}
}
