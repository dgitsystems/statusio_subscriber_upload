package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

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
	if len(os.Args) == 1 {
		log.Fatal("No filename for subscribers list given on the command line")
	}
	if len(os.Args) > 2 {
		log.Fatal("Too many command line parameters given.  This program only accepts a single parameter, the filename of the subscribers list to add")
	}
	fileName := os.Args[1]

	// Ensure file is present and readable
	info, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	if !info.Mode().IsRegular() {
		log.Fatal("File isn't a regular file")
	}

	// Read the file into memory
	tmpFile, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer tmpFile.Close()
	r := csv.NewReader(tmpFile)
	lines, err := r.ReadAll()
	if err != nil {
		return
	}
	addList := make(map[string]struct{})
	for _, field := range lines {
		addList[field[0]] = struct{}{}
	}

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
		fmt.Println()
	}

	// Remove any existing subscribers from the list of people to add
	for _, email := range subList.Result.Email {
		if _, ok := addList[email.Address]; ok {
			delete(addList, email.Address)
		}
	}

	if DEBUG {
		fmt.Println("Adding new subscribers...")
	}

	// Add the new subscribers
	for email := range addList {
		addInfo := statusio.Subscriber{
			StatuspageID: statusPageId,
			Method:       "email",
			Address:      email,
			Silent:       "1",
		}
		_, err := api.SubscriberAdd(addInfo)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v successfully subscribed\n", email)

		// Ensure there is at least a 1 second pause between subscriber add calls, so we don't hit rate limits
		// This is what the status.io docs say to do (really)
		time.Sleep(1100 * time.Millisecond) // Use 1.1 seconds, for that little bit extra safety
	}
}
