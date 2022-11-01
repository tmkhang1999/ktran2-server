package main

import (
	"github.com/gorilla/mux"
	"github.com/jamespearly/loggly"
	"log"
	"main.go/utils"
	"net/http"
	"strconv"
)

func main() {
	// Load variables
	config, err := utils.LoadConfig("./", "config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	// Instantiate the user variable including the loggly client
	user := utils.User{
		LogglyClient:   loggly.New("Weather-App"),
		DynamoDBClient: utils.CreateDynamoDBClient(config.Region),
		Config:         config,
	}

	// Set up mux router
	r := mux.NewRouter()
	r.HandleFunc("/ktran2/status", user.StatusHandler).Methods("GET")

	// Running
	srv := http.Server{
		Handler: r,
		Addr:    ":" + strconv.Itoa(config.Port),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
