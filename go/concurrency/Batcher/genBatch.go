package main

import (
	"fmt"
	"sync"
)

type Batcher struct {
	Input  chan int
	Output chan []int
}

func newBatcher() *Batcher {
	return &Batcher{
		Input:  make(chan int),
		Output: make(chan []int),
	}
}

func (batcher *Batcher) generateBatch(batchsize int, wg *sync.WaitGroup) {
	defer wg.Done()
	batch := []int{}
	for input := range batcher.Input { // listening to the input channel
		batch = append(batch, input)
		if len(batch) == batchsize {
			batcher.Output <- batch
			batch = []int{}
		}
		// time.Sleep(200 * time.Millisecond)
	}
	if len(batch) > 0 { // left out batch with bs < batchsize
		batcher.Output <- batch
		batch = []int{}
	}
	close(batcher.Output)

}

func usecaseBatcher(inp <-chan int, batcher *Batcher) {

	var wg sync.WaitGroup
	wg.Add(3)

	// Send inputs from users to the batcher
	go func() {
		for input := range inp {
			batcher.Input <- input
		}
		wg.Done()
		close(batcher.Input)
	}()

	bs := 5

	go batcher.generateBatch(bs, &wg)

	go func() {
		for op := range batcher.Output { // listen to output channel
			fmt.Printf("Output:%v\n", op)
		}
		wg.Done()
	}()
	wg.Wait()
}

func main() {

	batcher := newBatcher()
	userInp := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)
	go usecaseBatcher(userInp, batcher)

	// Generate dynamic input
	go sendInput([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, userInp, &wg)
	// go sendInput([]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 110, 111}, userInp)
	wg.Wait()
}

func sendInput(values []int, usersInput chan<- int, wg *sync.WaitGroup) {
	defer close(usersInput)
	defer wg.Done()
	for _, value := range values {
		usersInput <- value
	}
}
