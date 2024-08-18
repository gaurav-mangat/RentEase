package Extras

import (
	"RentEase/Internals/Authentication"
	"fmt"
)

func LandlordSection() {

	for {
		fmt.Println("\n\n\n----------------------------------------------")
		fmt.Println("             Landlord Section")
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
			// Call the Login method here
			//Login()
		case 2:
			// Call the SignUp method on signupHandler
			Authentication.SignUp()
		case 3:
			fmt.Println("Exiting the application. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
