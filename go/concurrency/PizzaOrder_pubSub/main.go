package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

type Producer struct {
	data chan PizzaOrder // pizza order channel
	quit chan bool       // quit channel to stop producing -> close the input stream
}

type PizzaOrder struct {
	orderNum int
	message  string
	success  bool
}

const (
	nPizza = 10
)

var (
	pizzasFailed = 0
	pizzasMade   = 0
	totalPizza   = 0
)

func makePizzaOrder(pizzaNum int) *PizzaOrder {
	var msg string
	success := false
	pizzaNum++
	if pizzaNum > nPizza {
		fmt.Println("Cant take this order. Pizza order exceeded")
	}
	timeToMake := rand.Intn(5)

	time.Sleep(1 * time.Second)
	fmt.Printf("Received order #%d!\n", pizzaNum)

	rnd := rand.Intn(12) // create random number generator to create multiple scenarios

	fmt.Printf("Started making pizza for order: #%d. It will take %d seconds\n", pizzaNum, timeToMake)

	if rnd <= 2 {
		pizzasFailed++
		msg = fmt.Sprintf("** Not having sufficient ingredients to make this order :#%d.** \n", pizzaNum)
	} else if rnd <= 4 {
		pizzasFailed++
		msg = fmt.Sprintf("** Currently not serving this kind of pizza for the order :#%d.** \n", pizzaNum)
	} else {
		pizzasMade++
		success = true
		msg = fmt.Sprintf("Pizza is ready for order :#%d.\n", pizzaNum)
	}
	totalPizza++

	return &PizzaOrder{
		orderNum: pizzaNum,
		message:  msg,
		success:  success,
	}
}

func pizzaProducer(producer *Producer) {
	// create pizza order
	pizzaNum := 0
	for { // to quit this loop has to use switch statement, push the data to the producer channel if want to quit close everything
		currentPizza := makePizzaOrder(pizzaNum) // count
		if currentPizza != nil {
			pizzaNum = currentPizza.orderNum
			select {
			case producer.data <- *currentPizza: // push pizza to this producer channel
			case <-producer.quit:
				close(producer.data)
				return
			}
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())

	color.Cyan("Pizza Town is open for business!")

	// create job to produce pizza
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan bool),
	}

	go pizzaProducer(pizzaJob) // run in the background

	for pizzaData := range pizzaJob.data {
		// read from the input streams
		if pizzaData.orderNum <= nPizza {
			if pizzaData.success {
				color.Green(pizzaData.message)
				color.Green("Order :#%d is out for delivery!", pizzaData.orderNum)
			} else {
				color.Red(pizzaData.message)
				color.Red("Order :#%d failed", pizzaData.orderNum)
			}
		} else {
			// order exceeded so close the business
			color.Cyan("Pizza Town is close for the business!")
			pizzaJob.quit <- true
		}
	}

	color.Cyan("Pizza order made: %d\npizza order Failed: %d\nTotal pizza order: %d\n", pizzasMade, pizzasFailed, totalPizza)
}
