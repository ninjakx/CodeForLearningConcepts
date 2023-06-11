package main

import (
	"fmt"
	"time"

	sm "pubsub"
)

func producer(stream *sm.Stream) (tweets []*sm.Tweet) {
	for {
		tweet, err := stream.Next()
		if err == sm.ErrEOF {
			return tweets
		}
		tweets = append(tweets, tweet)
	}
}

func consumer(tweets []*sm.Tweet) {
	for _, tweet := range tweets {
		if tweet.IsTopic1() {
			fmt.Println(tweet.Username, "tweet about topic-1")
		} else {
			fmt.Println(tweet.Username, "tweet about topic-2")
		}
	}
}

func main() {
	start := time.Now()
	stream := sm.GetMockStream()
	tweets := producer(stream)

	consumer(tweets)
	fmt.Printf("Process took %s\n", time.Since(start))
}
