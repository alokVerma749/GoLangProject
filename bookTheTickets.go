package main

import "fmt"

func bookTheTickets() {
	var availability uint = 50
	var total uint = 50
	var userFirstName string
	var userLastName string
	var ticket uint
	var booking []string
	fmt.Printf("Total tickets: %v\n", total)
	fmt.Printf("Available tickets: %v\n", availability)
	for i := availability; int(availability) >= 1; i++ {
		fmt.Println("Enter your first name")
		fmt.Scan(&userFirstName)
		fmt.Println("Enter your last name")
		fmt.Scan(&userLastName)
		fmt.Println("Enter the number of ticket to be booked")
		fmt.Scan(&ticket)
		// check if the ordering tickets are available or not
		if ticket <= availability {
			availability = availability - ticket
			fmt.Printf("%v your %v tickets are booked\n", userFirstName, ticket)
			booking = append(booking, userFirstName+" "+userLastName)
			fmt.Println("Press 0 to exit and 1 to continue")
			var choice uint
			fmt.Scan(&choice)
			if choice == 1 {
				continue
			} else if choice == 0 {
				break
			} else {
				fmt.Println("Enter the correct choice")
			}
		} else {
			fmt.Println("This much tickets are not available")
		}
	}
}
