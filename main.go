package main

import "fmt"

func main() {
	fmt.Println("Welcome to my ticet booking app")
	fmt.Println("Choose 1 to book the tickets ")
	fmt.Println("Choose 2 to check the ticket availability")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		bookTheTickets()
	case 0:
		bookTheTickets()
	}
}
