package main

import "fmt"

func main() {
	var num int

	fmt.Print("Enter any positive number to find Even or Odd: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("Number %d is EVEN\n", num)
	} else {
		fmt.Printf("Number %d is ODD\n", num)
	}
}
