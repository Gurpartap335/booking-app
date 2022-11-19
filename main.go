package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Conference Tickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v our conference booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you tickets here to attend")

	var bookings [50]string
	bookings[0] = "Nana"
	bookings[1] = "Eren"
	

	var firstName string
	var lastName string
	var emailAddress string
	var userTickets int


	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address")
	fmt.Scan(&emailAddress)

	fmt.Println("Please enter number of tickets")
	fmt.Scan(&userTickets) 

	remainingTickets -= userTickets;
	
	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
