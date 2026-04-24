package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, World!")
	time.Sleep(10 * time.Second)

	ch := make(chan string)

	go func() {
		ch <- "Hello from goroutine!"
	}()

	message := <-ch
	fmt.Println(message)
}
