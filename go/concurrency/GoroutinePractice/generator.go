package main

import "fmt"

func main() {
	generator := func(done <-chan bool, integers ...int) <-chan int {
		inputStream := make(chan int)

		go func() {
			defer close(inputStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case inputStream <- i:

				}
			}
		}()
		return inputStream
	}

	multiply := func(done <-chan bool, inputStream <-chan int, multiplier int) <-chan int {

		multiplyStream := make(chan int)
		go func() {
			defer close(multiplyStream)
			for i := range inputStream {
				select {
				case <-done:
					return
				case multiplyStream <- i * multiplier:
				}
			}
		}()
		return multiplyStream
	}

	done := make(chan bool)
	defer close(done)

	instream := generator(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	pipeline := multiply(done, instream, 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}
