package main

import (
	"LearnCoding/Golang/P1/hello"
	"fmt"
)

func main() {
	hello.Hello()
	fmt.Println("Printing hello world five times")
	for i := 0; i < 5; i++ {
		hello.Hello()
	}
}
