package utils

import (
	"fmt"
	"net/mail"
	"sync"
	"time"
)

var Wg = sync.WaitGroup{}

type UserData struct  {
	FirstName string
	LastName string
	Email string
	NumberOfTickets uint
}
func IsEmailValid(email string) bool {
    _, err := mail.ParseAddress(email)
	fmt.Printf("Parse %v email %v \n",email,err)
    return err == nil
}

func ValidateFields(userData UserData, remainingTickets uint)  bool {
	isValidName := len(userData.FirstName) >= 2 && len(userData.LastName) >= 2
	isValidEmail :=   IsEmailValid(userData.Email)
 
	isValidTickets := userData.NumberOfTickets > 0  
	shouldReturn := true
	if(!isValidName){
		fmt.Println("Invalid Name, please write 2 or more characters for First Name and Last Name")
		shouldReturn = false
	}

	if(!isValidEmail){
		fmt.Println("Invalid Email")
		shouldReturn = false
	}
	if(!isValidTickets ){
		fmt.Println("Invalid Ticket number, please write a number greater than 0")
		shouldReturn = false
	}

	if(userData.NumberOfTickets > remainingTickets ){
		fmt.Printf("We only have %v tickets remaining, so you cant book %v tickets\n",remainingTickets,userData.NumberOfTickets)
		shouldReturn =  false
	}
	return shouldReturn
}

func SendTicket(userData UserData){
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userData.NumberOfTickets,userData.FirstName,userData.LastName)
	fmt.Println("###################")
	fmt.Printf("Sending Ticket %v to %v\n",ticket,userData.Email)
	fmt.Println("###################")
	Wg.Done();
}