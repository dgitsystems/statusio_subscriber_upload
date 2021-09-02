package main

import (
	"fmt"
	"log"
	"os"

	"github.com/statusio/statusio-go"
)

var (
	apiId, apiKey, statusPageId string
	DEBUG                       = false
)

func main() {
	// Retrieve values from environment variables
	apiId = os.Getenv("API_ID")
	if apiId == "" {
		log.Fatal("API_ID environment variable missing")
	}
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY environment variable missing")
	}
	statusPageId = os.Getenv("STATUS_PAGE_ID")
	if statusPageId == "" {
		log.Fatal("STATUS_PAGE_ID environment variable missing")
	}

	// If the DEBUG variable is set to "true", then enable debugging
	dbgVar := os.Getenv("DEBUG")
	if dbgVar == "true" {
		DEBUG = true
	}

	// Check for filename being given on command line
	// TODO

	// Ensure file is present and readable
	// TODO

	// Read the file into memory
	// TODO

	// Establish connection to status.io
	api := statusio.NewStatusioApi(apiId, apiKey)

	// Retrieve the current (email) subscriber list
	subList, err := api.SubscriberList(statusPageId)
	if err != nil {
		log.Fatal(err)
	}

	if DEBUG {
		fmt.Println("Current subscriber list")
		fmt.Println("***********************")
		for _, sub := range subList.Result.Email {
			fmt.Println(sub.Address)
		}
	}

	// Remove any existing subscribers from the list of people to add
	// TODO

	// Loop through the list of new subscribers to add, adding them
	// TODO

	// Ensure there is at least a 1 second pause between subscriber add calls, so we don't hit rate limits
	// This is what the status.io docs say to do (really)
	// TODO
}
