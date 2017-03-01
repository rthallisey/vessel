package app

import (
	log "github.com/vessel/pkg/logger"
	"time"
)

func Run() *int {
	masterswitch := make(chan bool)

	for {
		log.Info("Starting main loop")
		a := new(int)
		time.Sleep(time.Duration(20) * time.Second)
		masterswitch <- true
		return a
	}
}
