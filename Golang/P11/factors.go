package main

import "fmt"

func main() {
	var fnum int

	fmt.Print("Enter any number find Factors: ")
	fmt.Scanln(&fnum)

	factor(fnum)

}

func factor(num int) {
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
