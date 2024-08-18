package main

import (
	"RentEase/ui"

	//"RentEase/Internals/Landlord"
	"fmt"
	"github.com/fatih/color"
)

func main() {

	// Print the startup message
	color.Red("**************************************")
	color.Red("*                                    *")
	color.Red("*           RentEase                 *")
	color.Red("*                                    *")
	color.Red("**************************************")

	color.Red("Welcome to RentEase, your CLI-based property rental application.")
	color.Red("Manage your properties, handle rentals, and keep track of everything in one place.")
	color.Red("Enjoy a seamless and efficient rental experience with our easy-to-use interface.")
	color.Red("For more information, type 'help' or 'man' for the user manual.")

	for {
		// Display menu options
		fmt.Println("\n\n\nPress 1 to go to Landlord section")
		fmt.Println("Press 2 to go to Tenant section")
		fmt.Println("Press 3 to go to Admin section")
		fmt.Println("Press 4 to exit!!")

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
			fmt.Println("You selected the Landlord section.")
			ui.UserSection("Landlord")
			// Add code to handle Landlord section here
		case 2:
			fmt.Println("You selected the Tenant section.")
			// Add code to handle Tenant section here
			ui.UserSection("Tenant")

		case 3:
			fmt.Println("You selected the Tenant section.")
			ui.UserSection("Admin")

		case 4:
			fmt.Println("Exiting the application. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
