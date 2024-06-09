package main

import "fmt"


func main() {
	canal := make(chan string)

	go func() {
		canal <- "some message"
	}()

	msg := <-canal
	fmt.Println((msg))
}