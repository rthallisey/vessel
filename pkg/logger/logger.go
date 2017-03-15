package logger

import (
	"io"
	"log"
	"os"
)

var Error *log.Logger
var Warning *log.Logger
var Info *log.Logger

func init() {

	errorHandle := io.Writer(os.Stderr)
	Err := log.New(errorHandle,
		"ERROR: ",
		log.Ltime|log.Lshortfile)

	warningHandle := io.Writer(os.Stdout)
	Warn := log.New(warningHandle,
		"WARNING: ",
		log.Lshortfile)

	infoHandle := io.Writer(os.Stdout)
	Information := log.New(infoHandle,
		"INFO: ",
		log.Lshortfile)

	Error = Err
	Warning = Warn
	Info = Information
}

func NewLogger() (*log.Logger, *log.Logger, *log.Logger) {
	return Error, Warning, Info
}
