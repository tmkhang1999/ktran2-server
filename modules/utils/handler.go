package utils

import (
	"encoding/json"
	"github.com/jamespearly/loggly"
	"main.go/structs"
	"net/http"
	"time"
)

type User struct {
	LogglyClient *loggly.ClientType
}

func (user *User) StatusHandler(w http.ResponseWriter, req *http.Request) {
	// Create a message to send to Loggly
	LogglyMessage := structs.LogglyMessage{
		StatusCode:  http.StatusOK,
		MethodType:  req.Method,
		SourceIp:    req.Host,
		RequestPath: req.URL.Path,
	}

	// Marshal the message struct
	message, marshalErr := json.Marshal(LogglyMessage)
	HandleException(marshalErr, user.LogglyClient, "Successfully marshal the loggly message")

	// Send the message to loggly
	SendingLoggy(user.LogglyClient, "info", string(message))

	// Create an HTTP response
	w.Header().Set("Content-Type", "application/json")
	resp := structs.Response{
		SystemTime: time.Now(),
		Status:     http.StatusOK,
	}

	// Marshal the HTTP response struct
	jsonResp, httpErr := json.Marshal(resp)
	HandleException(httpErr, user.LogglyClient, "Successfully marshal the http response")

	// Send the response to client
	_, writeErr := w.Write(jsonResp)
	HandleException(writeErr, user.LogglyClient, "Successfully send the http response")
}
