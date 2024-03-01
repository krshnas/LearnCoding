package main

import (
	add "LearnCoding/Golang/P2/addition"
	"fmt"
)

func main() {
	result := add.AddTwoNumbers(10, 20)
	fmt.Println("Result: ", result)

	var num1, num2 int
	fmt.Println("Enter any two number: ")
	fmt.Scan(&num1, &num2)
	fmt.Printf("num1 type: %T num2 type: %T\nnum1: %v num2: %v\n", num1, num2, num1, num2)

	result = add.AddTwoNumbers(num1, num2)
	fmt.Println("Result: ", result)
}
