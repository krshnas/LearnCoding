/*
Write a Go Program to calculate Compound Interest. This golang program allows the user to enter the principal amount, totals years,
and interest rates and then find the Compound Interest.
Future Compound Interest =  principal amount * (1 + interest rates)years
Compound Interest =  Future Compound Interest – principal amount
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var ci, pa, tys, ir, fci float64

	fmt.Println("Enter Principal Amount: ")
	fmt.Scanln(&pa)

	fmt.Printf("Enter Interest Rates: ")
	fmt.Scanln(&ir)

	fmt.Printf("Enter the Total Years: ")
	fmt.Scanln(&tys)

	fci = pa * (math.Pow((1 + ir/100), tys))
	ci = fci - pa

	fmt.Println("Compound Interest : ", ci)
	fmt.Println("Future Compound Interest : ", fci)
}
