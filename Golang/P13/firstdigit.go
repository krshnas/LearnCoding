package main

import "fmt"

func main() {
	var num, firstDigit int

	fmt.Print("Enter any number to return first digit: ")
	fmt.Scanln(&num)

	firstDigit = findFirstDigit(num)
	fmt.Println(firstDigit)
}

func findFirstDigit(num int) int {
	for num >= 10 {
		num /= 10
	}
	return num
}
