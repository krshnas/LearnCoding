package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: occtl --add <num1> <num2>")
		return
	}

	if os.Args[1] != "--add" {
		fmt.Println("Unknown operation. Use --add.")
		return
	}

	num1, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid number:", os.Args[2])
		return
	}

	num2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid number:", os.Args[3])
		return
	}

	result := num1 + num2
	fmt.Println("Result:", result)
}
