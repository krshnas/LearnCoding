package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter any Number: ")
	fmt.Scan(&num)
	result := findCube(num)
	fmt.Printf("Cube of %v is %v", num, result)
}

func findCube(num int) int {
	return num * num * num
}
