package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets = 50
var conferenceName = "Go Conference"
var remainingTickets = 50
var bookings = make([]UserData, 0)

type UserData struct {
		firstName 		string
		lastName		string
		email			string
		numberOfTickets	int
}

var wg = sync.WaitGroup{}

func main() {

		greetUsers()
		
		for {
			firstName, lastName, email, userTickets := getUserInput()
			isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

			if isValidName && isValidEmail && isValidTicketNumber {

					bookTicket(userTickets, firstName, lastName, email)

					wg.Add(1)
					go sendTicket(userTickets, firstName, lastName, email)

					firstNames := getFirstNames()
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
		wg.Wait()
		
}


func greetUsers() {
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
	 	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
		firstNames := []string{}
		for _, booking := range bookings {
			firstNames = append(firstNames, booking.firstName)
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


func bookTicket(userTickets int, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets) 
}


func sendTicket(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("################")
	wg.Done()
}