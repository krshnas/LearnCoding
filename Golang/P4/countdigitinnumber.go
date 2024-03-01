package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter any Number: ")
	fmt.Scanln(&num)
	count := 0
	fmt.Printf("Digits in Number %v is ", num)

	for num > 0 {
		count++
		num /= 10
	}

	fmt.Printf("%v", count)
}
