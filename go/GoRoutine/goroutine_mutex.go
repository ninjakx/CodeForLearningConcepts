package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type mutex sync.Mutex

func worker(ctx context.Context, mu *sync.Mutex) {
	for {
		deadline, _ := ctx.Deadline()
		fmt.Println(time.Until(deadline))

		select {
		case <-ctx.Done():
			return
		default:
			mu.Lock()
			// Critical section of code
			fmt.Println("==")
			mu.Unlock()
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	mu := sync.Mutex{}
	go worker(ctx, &mu)
	defer cancel()
	time.Sleep(10 * time.Second)
}
