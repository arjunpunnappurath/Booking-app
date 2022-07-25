package main

import (
	"fmt"
	"strings"
)

func main() {
	//var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	//arrays in go has fixed lengths
	//var bookings [50]string
	//slices
	bookings := []string{}
	//var remainingTIckets uint = 50 //uint will make sure the variable does not take a negative value

	//Print the types
	fmt.Printf("conferenceTickets is of the type %T, remaingingTickets is of %T and conferenceName is of %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Hello to our", conferenceName, "Booking Application")
	fmt.Printf("Welcome to  %v booking application\n", conferenceName)
	//fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here...")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		//ask user for the name
		fmt.Println("Enter the FirstName")
		fmt.Scan(&firstName)
		fmt.Println("Enter the last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter the email address")
		fmt.Scan(&email)
		fmt.Println("Enter number of tickets")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v%v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				//strings.fields() splits srings with white space as seperators
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("These first names of the bookings are %v\n", firstNames)
			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				//end program
				fmt.Println("Our conference is booked out. Please come over next year...")
				break
			}
		} else {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			//break
		}
	}
}
