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

func example1() {
	go sayHello()
	go sayHi()

	time.Sleep(800 * time.Millisecond)
}

func task(id int) {
	fmt.Println("Completed task", id)
}

func main() {
	for i := range 10 {
		//go task(i)
		go func(i int) {
			fmt.Println("Completed task", i)
		}(i)
	}

	time.Sleep(time.Second * 2)
}
