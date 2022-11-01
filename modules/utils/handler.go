package utils

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jamespearly/loggly"
	"log"
	"main.go/structs"
	"net/http"
)

type User struct {
	LogglyClient   *loggly.ClientType
	DynamoDBClient *dynamodb.DynamoDB
	Config         structs.Config
}

func updateRequestRecord(req *http.Request, logglyClient *loggly.ClientType) {
	// Create a message to send to Loggly
	LogglyMessage := structs.LogglyMessage{
		StatusCode:  http.StatusOK,
		MethodType:  req.Method,
		SourceIp:    req.Host,
		RequestPath: req.URL.Path,
	}

	// Marshal the message struct
	message, marshalErr := json.Marshal(LogglyMessage)
	if marshalErr != nil {
		log.Fatalln(marshalErr)
	}

	// Send the message to loggly
	SendingLoggy(logglyClient, "info", string(message))

}

func sendHTTPResponse(w http.ResponseWriter, logglyClient *loggly.ClientType, jsonResp []byte) {
	w.Header().Set("Content-Type", "application/json")
	_, writeErr := w.Write(jsonResp)
	HandleException(writeErr, logglyClient, "Successfully send the http response")
}

func (user *User) StatusHandler(w http.ResponseWriter, req *http.Request) {
	// Send a record of request to Loggly
	updateRequestRecord(req, user.LogglyClient)

	// Create the HTTP 'status' response
	httpStatusMessage := GetStatus(user.DynamoDBClient, user.Config.TableName)

	// Marshal the HTTP response struct
	jsonResp, httpErr := json.Marshal(httpStatusMessage)
	HandleException(httpErr, user.LogglyClient, "Successfully marshal the http response")

	// Send the response to client
	sendHTTPResponse(w, user.LogglyClient, jsonResp)
}

func (user *User) AllHandler(w http.ResponseWriter, req *http.Request) {
	// Send a record of request to Loggly
	updateRequestRecord(req, user.LogglyClient)

	// Create an HTTP 'all' response
	httpAllMessage := GetAll(user.DynamoDBClient, user.Config.TableName)

	// Marshal the HTTP response struct
	jsonResp, httpErr := json.Marshal(httpAllMessage)
	HandleException(httpErr, user.LogglyClient, "Successfully marshal the http response")

	// Send the response to client
	sendHTTPResponse(w, user.LogglyClient, jsonResp)
}
