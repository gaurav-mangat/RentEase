package Functionality

import (
	"RentEase/Internals/Landlord"
	"fmt"
)

func LandlordDashboard() {
	for {

		fmt.Println("\n\n\n----------------------------------------------")
		fmt.Println("             Landlord Dashboard")
		fmt.Println("----------------------------------------------")
		fmt.Println("\nPress 1 to login\nPress 2 to SignUp\nPress 3 to return")

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
			Landlord.Login()
		case 2:
			Landlord.SignUp()
		case 3:
			fmt.Println("Exiting the application. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
