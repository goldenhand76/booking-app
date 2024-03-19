package main

import (
	"booking-app/helper"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	conferenceName := "Go Conference" // := only works with variable not constant!
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	greatUsers(conferenceName)

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Println("we have total of", conferenceTickets, "Tickets and", remainingTickets, "Are still available")
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		var userData = make(map[string]string)
		// var userData = UserData{
		// 	firstName: firstName,
		// 	lastName: lastName,
		// 	email: email,
		// 	numberOfTickets: userTickets,
		// }
		type UserData struct {
			firstName       string
			lastName        string
			email           string
			numberOfTickets uint
		}

		helper.Validatereza("hi")
		// ask user for their name
		fmt.Println("Enter your last name:")
		fmt.Scan(&firstName)
		userData["firstName"] = firstName

		fmt.Println("Enter your first name:")
		fmt.Scan(&lastName)
		userData["lastName"] = lastName

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		fmt.Println("Enter your Number of tickets:")
		fmt.Scan(&userTickets)
		userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets.\n", firstName, lastName, userTickets)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of booking are: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("email address you entered is too short.")
			}
			if !isValidTicketNumber {
				fmt.Println("Ticket Number you entered is too short.")
			}
		}
	}
}

func greatUsers(confName string) {
	fmt.Printf("Welcome to %v booking application conference", confName)
}
