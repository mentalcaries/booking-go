package main

import (
	"fmt"
	"strings"
	"booking-app/helper"
)

var conferenceName = "GoConf"
const ticketNumber uint = 50
var remainingTickets uint = 50
var bookings = [] string{}

func main() {
	

	greetUsers()

	

	for remainingTickets > 0 && len(bookings) < 50 {
			

			firstName, lastName, email, userTickets := getUserInput()

			isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, uint(userTickets), remainingTickets)

			bookings = append(bookings, firstName + " " + lastName)


			if isValidName && isValidEmail && isValidTicketNumber {

				bookTicket(userTickets, firstName, lastName, email)
				
				firstNames := getFirstNames()

				fmt.Printf("These are the people with bookings %v\n", firstNames)


				if remainingTickets == 0 {
					fmt.Println("All tickets are sold out :( . Please try again next time")
					break
				}
			} else {
				if !isValidName{

					fmt.Println("First or last name too short, try again")
				}
				if !isValidEmail{

					fmt.Println("Invalid email, try again")
				}
				if !isValidTicketNumber{

					fmt.Println("Please enter a number greater than 0")
				}
				continue
			}

			
	}

}

func greetUsers(){
	fmt.Printf("Welcoming to %v\n", conferenceName, )
	fmt.Printf("Total number of tickets: %v. We have %v remaining. \n", ticketNumber, remainingTickets)
	fmt.Println("Give us your money, get a ticket")
}

func getFirstNames() []string {
	firstNames := []string{}
				for _ , booking := range bookings {
					var firstName = strings.Fields(booking)[0]
					firstNames = append(firstNames, firstName)
				}
	return firstNames
}



func getUserInput() (string, string, string, uint){

			var firstName string
			var lastName string
			var email string
			var userTickets uint

	fmt.Println("Enter your first name:")
			fmt.Scan(&firstName)
			
			fmt.Println("Enter your last name:")
			fmt.Scan(&lastName)
			
			fmt.Println("Enter your email:")
			fmt.Scan(&email)

			fmt.Println("Enter the number of tickets you want:")
			fmt.Scan(&userTickets)

			return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

				fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive your confirmation at %v \n", firstName, lastName, userTickets, email)
				fmt.Printf("%v tickets remain\n", remainingTickets)

}