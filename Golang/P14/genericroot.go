package main

import "fmt"

func main() {
	var num, sum int

	fmt.Print("Enter any Number to find Generic Root: ")
	fmt.Scanln(&num)

	for num >= 10 {
		sum = digitSum(num)
		if sum >= 10 {
			num = sum
		} else {
			fmt.Println("The Generic Root of this Number is ", sum)
			break
		}
	}
}

func digitSum(num int) int {
	var sum, rem int
	for sum = 0; num > 0; num = num / 10 {
		rem = num % 10
		sum += rem
	}
	return sum
}
