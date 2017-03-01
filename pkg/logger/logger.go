package test

import (
    "io"
    "log"
    "os"
)


func init() {
     NewLogger()
}

type Logger struct {
    Info *log.Logger
    Warning *log.Logger
    Error *log.Logger
}

func NewLogger() *Logger {
    infoHandle := io.Writer(os.Stdout)
    warningHandle := io.Writer(os.Stdout)
    errorHandle := io.Writer(os.Stderr)

    info := log.New(infoHandle,
        "INFO: ",
        log.Lshortfile)
    warning := log.New(warningHandle,
        "WARNING: ",
        log.Lshortfile)
    error := log.New(errorHandle,
        "Error: ",
        log.Ltime|log.Lshortfile)

    return &Logger{Info: info, Warning: warning, Error: error}
}
