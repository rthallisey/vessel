package app

import (
	l "github.com/vessel/pkg/logger"
	"time"
)

var logger = l.NewLogger()
var Info = logger.Info
var Warning = logger.Warning
var Error = logger.Error

func Run() *int {
	masterswitch := make(chan bool)

	for {
		Info.Print("Starting main loop")
		a := new(int)
		time.Sleep(time.Duration(20) * time.Second)
		masterswitch <- true
		return a
	}
}
