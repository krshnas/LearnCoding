package main

import "fmt"

func main() {

	var basicSal, hra, da, grossSal float64

	fmt.Print("Enter the Employee Basic Salary = ")
	fmt.Scanln(&basicSal)

	fmt.Print("Enter the Employee HRA Amount = ")
	fmt.Scanln(&hra)

	fmt.Print("Enter the Employee DA Amount = ")
	fmt.Scanln(&da)

	grossSal = basicSal + hra + da
	fmt.Println("The Gross Salary of this Employee = ", grossSal)
}
