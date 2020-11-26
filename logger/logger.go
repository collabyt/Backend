package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	Warning *log.Logger
	Error   *log.Logger
	Info    *log.Logger
)

func Setup() {
	currDateTime := time.Now().Format(time.RFC3339)
	newLogFile := fmt.Sprintf("logs/%s.txt", currDateTime)
	lf, err := os.OpenFile(newLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	Info = log.New(lf, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(lf, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(lf, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
