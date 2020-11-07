package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/kylelemons/godebug/pretty"
	"github.com/nathanperkins/reddit-experiment/src/reddit"
)

func main() {
	flag.Parse()
	redisAddr, ok := os.LookupEnv("REDIS_ADDR")
	if !ok {
		log.Fatalf("env var REDIS_ADDR is required")
	}

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Query redis
		var post reddit.Post
		val, err := rdb.Get(ctx, "top-post").Result()
		if err == redis.Nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "top post has not been set yet")
			return
		} else if err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "failed to contact db")
			return
		}

		// Decode response
		if err := post.UnmarshalBinary([]byte(val)); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "failed to decode from db")
			return
		}

		fmt.Fprintf(w, "Top post: %v\n", pretty.Sprint(post))
	})

	log.Fatal(http.ListenAndServe(":80", nil))
}
