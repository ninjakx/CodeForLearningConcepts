package main

import (
	"fmt"
	"sync"
)

var once sync.Once

type singleton struct {
}

var instance *singleton

func getInstance() *singleton {
	if instance == nil {
		once.Do(
			func() {
				fmt.Println("Creating instance...")
				instance = &singleton{}
			})
	} else {
		fmt.Println("Singleton already exists")
	}
	return instance
}

func main() {
	for i := 0; i < 10; i++ {
		getInstance()
	}
}
