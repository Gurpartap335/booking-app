package main

import (
	"fmt"
	"strings"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = []string{}

func main() {

	fmt.Printf("Welcome to %v our conference booking application", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get you tickets here to attend")

	bookings := []string{}
	
	fmt.Println(bookings)
	
	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

		   fmt.Printf("The first names of bookings are: %v\n", firstNames)

		   if remainingTickets == 0 {
			   // end program
			   fmt.Println("Our conference is booked out. Come back next year.")
			   break
		    }
	   } else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			}
	   }
	
    }
}


func greetUsers() {
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
	 	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
		firstNames := []string{}
		for _, booking := range bookings {
			var  names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		return firstNames
}

func getUserInput() (string, string, string, int) {
		var firstName string
		var lastName string
		var email string
		var userTickets int


		fmt.Println("Please enter your first name")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email address")
		fmt.Scan(&email)

		fmt.Println("Please enter number of tickets")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}


func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 3 && len(lastName) >= 3
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(remainingTickets int, userTickets int, bookings []string, firstName string, lastName string, email string, conference string) {
	remainingTickets -= userTickets;
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, emailAddress)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)

}