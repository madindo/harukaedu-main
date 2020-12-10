package logs

import (
	"log"
	"os"
)

func Logging(msgtype string, msg string) {
	file, err := os.OpenFile("logs/all.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print( " [ " + msgtype + " ] " + msg)
}