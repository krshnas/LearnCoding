package main

import "fmt"

func main() {
	var num1, num2, largest int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	largest = findLargest(num1, num2)
	if largest != 0 {
		fmt.Println("Largest number is ", largest)
	} else {
		fmt.Print("Both are Equal")
	}
}

func findLargest(num1, num2 int) int {
	if num1 == num2 {
		return 0
	} else if num1 > num2 {
		return num1
	}
	return num2
}
