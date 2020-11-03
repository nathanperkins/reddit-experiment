package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
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
	redisAddr, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		log.Fatalf("env var REDIS_ADDR is required")
	}

	ctx := context.Background()

	// Get clients for Reddit and Redis
	client, err := reddit.New(username, password, clientID, clientSecret)
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Query Reddit for top post in politics in last hour
	query := &url.Values{
		"t":     {"hour"},
		"limit": {"1"},
	}
	var resp reddit.Listing
	if err = client.Get("https://oauth.reddit.com/r/politics/top/", &resp, query); err != nil {
		panic(err)
	}

	// Update redis.
	post := resp.Data.Children[0].Data
	if _, err := rdb.Set(ctx, "top-post", post, time.Hour).Result(); err != nil {
		panic(err)
	}
	b, _ := json.Marshal(post)
	fmt.Printf("Updated top post: %v", string(b))

}
