package utils

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jamespearly/loggly"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"main.go/structs"
	"net/http"
	"strings"
	"time"
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

func (user *User) StatusHandler(w http.ResponseWriter, req *http.Request) {
	// Send a record of request to Loggly
	updateRequestRecord(req, user.LogglyClient)

	// Create the HTTP 'status' response
	w.Header().Set("Content-Type", "application/json")
	httpStatusMessage := GetStatus(user.DynamoDBClient, user.Config.TableName)

	// Marshal the HTTP response struct
	jsonResp, httpErr := json.Marshal(httpStatusMessage)
	HandleException(httpErr, user.LogglyClient, "Successfully marshal the http response")

	// Send the response to client
	_, writeErr := w.Write(jsonResp)
	HandleException(writeErr, user.LogglyClient, "Successfully send the http response")
}

func (user *User) AllHandler(w http.ResponseWriter, req *http.Request) {
	// Send a record of request to Loggly
	updateRequestRecord(req, user.LogglyClient)

	// Create an HTTP 'all' response
	w.Header().Set("Content-Type", "application/json")
	httpAllMessage := GetAll(user.DynamoDBClient, user.Config.TableName)

	// Marshal the HTTP response struct
	jsonResp, httpErr := json.Marshal(httpAllMessage)
	HandleException(httpErr, user.LogglyClient, "Successfully marshal the http response")

	// Send the response to client
	_, writeErr := w.Write(jsonResp)
	HandleException(writeErr, user.LogglyClient, "Successfully send the http response")
}

func (user *User) SearchHandler(w http.ResponseWriter, req *http.Request) {
	// Send a record of request to Loggly
	updateRequestRecord(req, user.LogglyClient)

	// Create an HTTP 'search' response
	w.Header().Set("Content-Type", "application/json")
	httpSearchMessage := user.getSearch(w, req)

	if httpSearchMessage != nil {
		// Marshal the HTTP response struct
		jsonResp, httpErr := json.Marshal(httpSearchMessage)
		HandleException(httpErr, user.LogglyClient, "Successfully marshal the http response")

		// Send the response to client
		_, writeErr := w.Write(jsonResp)
		HandleException(writeErr, user.LogglyClient, "Successfully send the http response")
	}
}

func (user *User) getSearch(w http.ResponseWriter, req *http.Request) []structs.AwsItem {
	// Initialize the policy to sanitize incoming query requests
	policy := bluemonday.StrictPolicy()

	// Sanitize the query
	query := req.URL.Query()
	city := cleanUpWords(policy.Sanitize(query.Get("city")))

	// If a query parameter is anything other than city, return 400.
	if city == "" || len(query) == 0 {
		user.sendBadRequest(w, 400, "Invalid query!")
		return nil
	}

	// Filter item results based on the query parameter
	var results []structs.AwsItem
	AwsItems := GetAll(user.DynamoDBClient, user.Config.TableName)
	for _, item := range AwsItems {
		if cleanUpWords(item.Name) == city {
			results = append(results, item)
		}
	}

	// If the search returns no results, return 404.
	if len(results) == 0 {
		user.sendBadRequest(w, 404, "Sorry! We could not find your city in our DynamoDB")
		return nil
	}

	return results
}

func (user *User) sendBadRequest(w http.ResponseWriter, statusCode int, msg string) {
	// Write status code in header
	w.WriteHeader(statusCode)

	// Create a http response
	resp := structs.Response{
		SystemTime: time.Now(),
		Error:      msg,
	}

	// Marshal the HTTP response struct
	jsonResp, httpErr := json.Marshal(resp)
	HandleException(httpErr, user.LogglyClient, "Successfully marshal the http response")

	// Send the response to client
	_, writeErr := w.Write(jsonResp)
	HandleException(writeErr, user.LogglyClient, "Successfully send the http response")
}

func cleanUpWords(input string) string {
	result := strings.ToLower(strings.ReplaceAll(input, " ", ""))
	return result
}
