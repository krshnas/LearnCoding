package main

import "fmt"

func main() {
	var units, surCharge, amount, totAmount float64

	fmt.Print("Enter the Consumed Units: ")
	fmt.Scanln(&units)

	if units < 50 {
		amount = units * 2.60
		surCharge = 25
	} else if units <= 100 {
		amount = 130 + ((units - 50) * 3.25)
		surCharge = 35

	} else if units <= 200 {
		amount = 130 + 162.50 + ((units - 100) * 5.26)
		surCharge = 45
	} else {
		amount = 130 + 162.50 + 526 + ((units - 200) * 7.75)
		surCharge = 55
	}
	totAmount = amount + surCharge
	fmt.Println("Electricity Bill:", totAmount)
}
