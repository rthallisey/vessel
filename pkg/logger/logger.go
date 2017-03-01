package logger

import (
	"io"
	"log"
	"os"
)

func Error(err string) {
	errorHandle := io.Writer(os.Stderr)

	Err := log.New(errorHandle,
		"ERROR: ",
		log.Ltime|log.Lshortfile)
	Err.Println(err)
}

func Warning(err string) {
	warningHandle := io.Writer(os.Stdout)

	Warn := log.New(warningHandle,
		"WARNING: ",
		log.Lshortfile)
	Warn.Println(err)
}

func Info(err string) {
	infoHandle := io.Writer(os.Stdout)

	Information := log.New(infoHandle,
		"INFO: ",
		log.Lshortfile)
	Information.Println(err)
}
