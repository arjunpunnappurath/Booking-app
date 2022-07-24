package main

import "fmt"

func main() {
	//var conferenceName = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	//arrays in go has fixed lengths
	var bookings [50]string
	//var remainingTIckets uint = 50 //uint will make sure the variable does not take a negative value

	//Print the types
	fmt.Printf("conferenceTickets is of the type %T, remaingingTickets is of %T and conferenceName is of %T\n", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Hello to our", conferenceName, "Booking Application")
	fmt.Printf("Welcome to  %v booking application\n", conferenceName)
	//fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are still available")
	fmt.Printf("We have a total of %v and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here...")

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

	fmt.Printf("Thank you %v%v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	//fmt.Println(remainingTickets)
	//fmt.Println(&remainingTickets)

	//userName = "TOM"
	//userTickets = 2

	remainingTickets = remainingTickets - userTickets
	bookings[0] = firstName + " " + lastName

	fmt.Printf("The whole array : %v\n", bookings)
	fmt.Printf("The first value : %v\n", bookings[0])
	fmt.Printf("Array type %T\n", bookings)
	fmt.Printf("Array length %v\n", len(bookings))

	fmt.Printf("%v tickets are remaining for %v", remainingTickets, conferenceName)
	//fmt.Printf("User %v booked %v tickets\n", firstName, userTickets)ß
}
