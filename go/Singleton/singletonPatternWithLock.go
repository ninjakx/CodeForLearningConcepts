package main

import (
	"fmt"
	"sync"
)

type singleton struct {
}

var lock = &sync.Mutex{}

var instance *singleton

func getInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("creating singleton instance...")
			instance = &singleton{}
		} else {
			fmt.Println("Singleton instance already created...")
		}
	} else {
		fmt.Println("Singleton instance already created.")
	}
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		getInstance()
	}
}
