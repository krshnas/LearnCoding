package main

import "fmt"

func main() {
	var num1, num2, num3, largest, largest2 int

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	fmt.Print("Enter third number: ")
	fmt.Scanln(&num3)

	largest = findLargest(num1, num2, num3)
	if largest != 0 {
		fmt.Println("Largest number is ", largest)
	} else {
		fmt.Print("Both are Equal")
	}
	largest2 = findLargest2(num1, num2, num3)
	fmt.Println("Largest number is ", largest2)
}

func findLargest(num1, num2, num3 int) int {
	if num1 > num2 && num1 > num3 {
		return num1
	} else if num2 > num1 && num2 > num3 {
		return num2
	} else if num3 > num1 && num3 > num2 {
		return num3
	}
	return 0
}

func findLargest2(num1, num2, num3 int) int {

	if num1-num2 > 0 && num1-num3 > 0 {
		return num1
	} else {
		if num2-num3 > 0 {
			return num2
		} else {
			return num3
		}
	}
}
