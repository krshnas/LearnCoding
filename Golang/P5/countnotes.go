package main

import "fmt"

func main() {
	var amount int
	fmt.Print("Enter amount: ")
	fmt.Scanln(&amount)
	CountNotes(amount)
}

func CountNotes(amount int) {
	notes := [8]int{500, 100, 50, 20, 10, 5, 2, 1}
	temp := amount

	for i := 0; i < 8; i++ {
		fmt.Println(notes[i], " Notes = ", temp/notes[i])
		temp %= notes[i]
	}
}
