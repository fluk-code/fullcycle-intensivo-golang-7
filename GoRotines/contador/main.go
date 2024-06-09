package main

import (
	"fmt"
	"time"
)


func contador(x int, name string ) {
	for i := 0; i < x; i++ {
		fmt.Print(name + " ")
		fmt.Println(i)
		time.Sleep((time.Second))
	}
}

func main() {
	go contador(10, "T1")
	contador(10, "T2")
}