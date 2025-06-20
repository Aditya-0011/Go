package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hello, again!")
}

func sayHi() {
	fmt.Println("Hi, there!")
}

func main() {
	go sayHello()
	go sayHi()

	time.Sleep(800 * time.Millisecond)
}
