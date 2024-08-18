package ui

import (
	"RentEase/Internals/Authentication"
	"fmt"
)

func UserSection(user string) {

	for {

		// Printing userSection based on the type of user
		fmt.Println("\n\n\n----------------------------------------------")
		if user == "Landlord" {
			fmt.Println("             Landlord Section")
		} else if user == "Tenant" {
			fmt.Println("             Tenant Section")
		} else {
			fmt.Println("             Admin Section")
		}
		fmt.Println("----------------------------------------------")

		fmt.Println("\nPress 1 to login")
		if user != "Admin" {
			fmt.Println("Press 2 to SignUp")

		}
		fmt.Println("Press 3 to return")

		// Read user input
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// Handle user choice using switch-case
		switch choice {
		case 1:
			// Call the Login method here
			Authentication.Login(user)
		case 2:
			// Call the SignUp
			if user == "Admin" {
				fmt.Println("Invalid choice. Please select a valid option.")
				return
			}
			Authentication.SignUp(user)
		case 3:
			fmt.Println("Exiting the application. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
