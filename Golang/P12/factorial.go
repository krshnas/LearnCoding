package main

import "fmt"

func main() {
	var fnum int

	fmt.Print("Enter any Number to find Factorial: ")
	fmt.Scanln(&fnum)

	factorial := factorial(fnum)
	fmt.Printf("Factorial of %d is %v", fnum, factorial)
}

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
