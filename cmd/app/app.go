package app

import (
       l "github.com/vessel/pkg/logger"
)

var logger = l.NewLogger()
var Info = *logger.Info
var Warning = *logger.Warning
var Error = *logger.Error

func Run () *int {
     for {
     	 Info.Print("Starting main loop")
	 t := new(int)
	 return t
     }
}
