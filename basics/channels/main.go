package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNumber(channel chan int) {
	for num := range channel {
		fmt.Println("Processing number:", num)
		time.Sleep(time.Second)
	}
}

func example1() {
	numberChannel := make(chan int)
	go processNumber(numberChannel)

	for {
		numberChannel <- rand.Intn(100)
	}
}

func sum(channel chan int, a int, b int) {
	result := a + b
	channel <- result
}

func example2() {
	result := make(chan int)

	go sum(result, 5, 10)

	res := <-result
	fmt.Println("Sum:", res)
}

func task(channel chan bool) {
	defer func() { channel <- true }()
	fmt.Println("processing...")
}

func example3() {
	taskChannel := make(chan bool)
	go task(taskChannel)
	<-taskChannel
}

func messageSender(channel <-chan string, done chan<- bool) { //channel is read-only, done is write-only
	defer func() { done <- true }()
	for message := range channel {
		fmt.Println("Sending message:", message)
		time.Sleep(time.Second)
	}
}

func example4() {
	messageChannel := make(chan string, 100)
	done := make(chan bool)

	go messageSender(messageChannel, done)

	for i := range 10 {
		messageChannel <- fmt.Sprintf("Message %d", i)
	}

	fmt.Println("Done sending")

	close(messageChannel)
	<-done
}

func example5() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() { chan1 <- 10 }()
	go func() { chan2 <- "Hello" }()

	for range 2 {
		select {
		case num := <-chan1:
			fmt.Println("Received from chan1:", num)
		case str := <-chan2:
			fmt.Println("Received from chan2:", str)
		}
	}
}

func main() {
	// example1()

	// example2()

	// example3()

	// example4()

	example5()
}
