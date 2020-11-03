package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/nathanperkins/reddit-experiment/src/reddit"
)

func main() {
	flag.Parse()
	username, ok := os.LookupEnv("REDDIT_USERNAME")
	if !ok {
		log.Fatalf("env var $REDDIT_USERNAME is required")
	}
	password, ok := os.LookupEnv("REDDIT_PASSWORD")
	if !ok {
		log.Fatalf("env var $REDDIT_PASSWORD is required")
	}
	clientID, ok := os.LookupEnv("REDDIT_CLIENT_ID")
	if !ok {
		log.Fatalf("env var REDDIT_CLIENT_ID is required")
	}
	clientSecret, ok := os.LookupEnv("REDDIT_CLIENT_SECRET")
	if !ok {
		log.Fatalf("env var REDDIT_CLIENT_SECRET is required")
	}

	client, err := reddit.New(username, password, clientID, clientSecret)
	if err != nil {
		panic(err)
	}
	query := &url.Values{
		"t":     {"hour"},
		"limit": {"1"},
	}
	type listing struct {
		Data struct {
			Children []struct {
				Data struct {
					Title     string
					Created   float32
					Permalink string
					Score     float32
				}
			}
		}
	}
	var resp listing
	if err = client.Get("https://oauth.reddit.com/r/politics/top/", &resp, query); err != nil {
		panic(err)
	}
	post := resp.Data.Children[0].Data
	fmt.Printf("Top post: %v\n", post)
}
