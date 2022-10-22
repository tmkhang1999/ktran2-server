package main

import (
	"github.com/gorilla/mux"
	"github.com/jamespearly/loggly"
	"log"
	"main.go/utils"
	"net/http"
)

func main() {
	// Load variables
	_, err := utils.LoadConfig("./", "config.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	// Instantiate the user variable including the loggly client
	user := utils.User{
		LogglyClient: loggly.New("Weather-App"),
	}

	// Set up mux router
	r := mux.NewRouter()
	r.HandleFunc("/ktran2/status", user.StatusHandler).Methods("GET")

	// Running
	err = http.ListenAndServe(":4000", r)
	if err != nil {
		return
	}
}
