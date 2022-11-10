package main

import (
	"booking-app/utils"
	"fmt"
)

var bookings = make([]utils.UserData,0)
var conferenceName = "Go Conference"
const conferenceTickets  = 50
var remainingTickets uint = conferenceTickets



func main() {

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)

	fmt.Println(conferenceName)

	fmt.Printf("Welcome to %v booking application\n",conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)

	fmt.Println("Get Em!")
	

	for (remainingTickets > 0){

		

		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their name
		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)
		
		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)
	
		fmt.Println("Please enter your email:")
		fmt.Scan(&email)
	
	
	
		fmt.Println("Please the number of tickets:")
		fmt.Scan(&userTickets)
		
	
		var userData = utils.UserData {
			FirstName: firstName,
			LastName: lastName,
			Email: email,
			NumberOfTickets: userTickets,
		}
	
	 
		if(!utils.ValidateFields(userData,remainingTickets)){
			continue;
		}
		remainingTickets =  remainingTickets - userTickets
	 
		bookings = append(bookings, userData)
	 
	
		fmt.Printf("Thank you %v %v for booking your %v ticket/s.\n You should recieve an email confirmation at %v\n",firstName,lastName,userTickets,email);
	
		fmt.Printf("There are %v remaining tickets for %v\n",remainingTickets,conferenceName)
	
		firstNames := []string{}
		for _, booking := range bookings{ 
			firstNames = append(firstNames, booking.FirstName)
		}

		fmt.Printf("These are all our bookings: %v\n",firstNames)
		utils.Wg.Add(1)
		go utils.SendTicket(userData)


		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			fmt.Println("Out conference is booked out. Come back next year!")
			break
		}
	}

	utils.Wg.Wait()
	
}
