package utils

import (
	"github.com/jamespearly/loggly"
	"log"
)

func SendingLoggy(client *loggly.ClientType, msgType string, msg string) {
	err := client.EchoSend(msgType, msg)
	if err != nil {
		log.Fatalln(err)
	}
}

func HandleException(e error, client *loggly.ClientType, successfulMsg string) {
	if e == nil {
		SendingLoggy(client, "info", successfulMsg)
	} else {
		SendingLoggy(client, "error", "Failed with error: "+e.Error())
	}
}
