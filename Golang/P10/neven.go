package main

import "fmt"

func main() {
	var lnum int

	fmt.Print("Enter the Range of Number to Print Even's 1 to ")
	fmt.Scanln(&lnum)

	for i := 2; i <= lnum; i += 2 {
		fmt.Print(i, "\t")
	}
}
