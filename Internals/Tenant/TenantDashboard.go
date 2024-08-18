package Tenant

import (
	"RentEase/models"
	"fmt"
)

// TenantDashboard function to handle user options
func TenantDashboard(activeUser models.User) {
	activeTenant := activeUser

	//  Function to load properties once

	for {
		// User Dashboard
		fmt.Println()
		fmt.Println("\033[1;36m---------------------------------------------\033[0m")      // Sky blue
		fmt.Println("\033[1;32m             TENANT DASHBOARD                     \033[0m") // Green
		fmt.Println("\033[1;36m---------------------------------------------\033[0m")      // Sky blue
		fmt.Println("\n\033[1;34m     1. Search for Property \033[0m")
		fmt.Println()
		fmt.Println("\033[1;34m     2. View Wishlist\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34m     3. Add Wishlist\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34m     3. Delete the Property\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34m     4. Respond to Inquiries\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34m     5. View your profile\033[0m")
		fmt.Println()
		fmt.Println("\033[1;34m     6. Exit\033[0m")
		fmt.Println()

		var choice int
		fmt.Print("     Enter your choice: ")
		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		switch choice {
		case 1:
			searchProperties()
		case 2:
			ViewWishlist(activeTenant.Username)
		case 3:
			var propertyId int
			fmt.Print("Enter the property Id: ")
			fmt.Scan(&propertyId)
			AddToWishlist(propertyId, activeTenant.Username)
		case 4:
			// respondToInquiries() // Placeholder function
		case 5:
			//viewProfile()
		case 6:
			fmt.Println("Logging out...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
