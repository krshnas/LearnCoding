package main

import "fmt"

var count int = 0

func digitCount(num int) int {
    if num > 0 {
        count = count + 1
        digitCount(num / 10)
    }
    return count
}

func main() {

    var num int

    fmt.Print("Enter any number to count digits = ")
    fmt.Scanln(&num)

    count = digitCount(num)
    fmt.Println("The total Number to Digits = ", count)
}