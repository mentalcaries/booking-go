package main

import (
	"fmt"
	"strings"
)


func main() {
	
	conferenceName := "GoConf"
	const ticketNumber = 50
	var remainingTickets = 50

	greetUsers(conferenceName, ticketNumber, uint(remainingTickets))

	
	var bookings [] string

	for remainingTickets > 0 && len(bookings) < 50 {
			var firstName string
			var lastName string
			var email string
			var userTickets int


			fmt.Println("Enter your first name:")
			fmt.Scan(&firstName)
			
			fmt.Println("Enter your last name:")
			fmt.Scan(&lastName)
			
			fmt.Println("Enter your email:")
			fmt.Scan(&email)

			fmt.Println("Enter the number of tickets you want:")
			fmt.Scan(&userTickets)

			isValidName := len(firstName) >=2 && len(lastName) >=2
			isValidEmail := strings.Contains(email, "@")
			isValidTicketNumber := userTickets >0 && userTickets <= remainingTickets


			bookings = append(bookings, firstName + " " + lastName)


			if isValidName && isValidEmail && isValidTicketNumber {

				remainingTickets -= userTickets

				fmt.Printf("Thank you %v %v for purchasing %v tickets. You will receive your confirmation at %v \n", firstName, lastName, userTickets, email)
				fmt.Printf("%v tickets remain\n", remainingTickets)

				printFirstNames((bookings))

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

				
			}
			
	}

}

func greetUsers(confName string, ticketNumber int, remainingTickets uint){
	fmt.Printf("Welcoming to %v\n", confName, )
	fmt.Printf("Total number of tickets: %v. We have %v remaining. \n", ticketNumber, remainingTickets)
	fmt.Println("Give us your money, get a ticket")
}

func printFirstNames(bookings []string){
	firstNames := []string{}
				for _ , booking := range bookings {
					var firstName = strings.Fields(booking)[0]
					firstNames = append(firstNames, firstName)
				}

				fmt.Printf("These are the people with bookings %v\n", firstNames)
}