package logger

import (
	"log"
	"os"
)

func Setup() {
	// TODO: MODIFY TO CREATE A NEW LOGFILE FOR EACH SERVER INITIALIZATION
	lf, err := os.OpenFile("logs/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	log.SetOutput(lf)
}
