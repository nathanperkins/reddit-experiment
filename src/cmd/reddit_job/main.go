package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nathanperkins/reddit-experiment/src/reddit"
)

func main() {
	flag.Parse()
	username, ok := os.LookupEnv("REDDIT_USERNAME")
	if !ok {
		log.Fatalf("--username is required")
	}
	password, ok := os.LookupEnv("REDDIT_PASSWORD")
	if !ok {
		log.Fatalf("--password is required")
	}
	clientID, ok := os.LookupEnv("REDDIT_CLIENT_ID")
	if !ok {
		log.Fatalf("--client_id is required")
	}
	clientSecret, ok := os.LookupEnv("REDDIT_CLIENT_SECRET")
	if !ok {
		log.Fatalf("--clientSecret is required")
	}

	client, err := reddit.New(username, password, clientID, clientSecret)
	if err != nil {
		panic(err)
	}
	var profileData struct {
		Name    string
		Created float32
	}
	if err = client.Get("https://oauth.reddit.com/api/v1/me", &profileData); err != nil {
		panic(err)
	}
	fmt.Printf("Data: %v\n", profileData)
}
