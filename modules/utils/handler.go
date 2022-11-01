package utils

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jamespearly/loggly"
	"main.go/structs"
	"net/http"
)

type User struct {
	LogglyClient   *loggly.ClientType
	DynamoDBClient *dynamodb.DynamoDB
	Config         structs.Config
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
	httpStatusMessage := GetStatus(user.DynamoDBClient, user.Config.TableName)

	// Marshal the HTTP response struct
	jsonResp, httpErr := json.Marshal(httpStatusMessage)
	HandleException(httpErr, user.LogglyClient, "Successfully marshal the http response")

	// Send the response to client
	_, writeErr := w.Write(jsonResp)
	HandleException(writeErr, user.LogglyClient, "Successfully send the http response")
}
