package Admin

import (
	"RentEase/utils"
	"fmt"
)

// ShowAdminDashboard displays the admin dashboard menu and handles user inputs.
func AdminDashboard() {
	for {
		fmt.Println()
		fmt.Println("\033[1;36m---------------------------------------------\033[0m")     // Sky blue
		fmt.Println("\033[1;32m             ADMIN DASHBOARD                     \033[0m") // Green
		fmt.Println("\033[1;36m---------------------------------------------\033[0m")
		fmt.Println("\n\n1. View All Users")
		fmt.Println("\n2. View Active Landlord Users")
		fmt.Println("\n3. View Inactive Landlord Users")
		fmt.Println("\n4. delete Landlord Users")
		fmt.Println("5. Exit")
		choice := utils.ReadInput("Enter your choice: ")

		switch choice {
		case "1":
			LoadAllUsers()
		case "2":
			fmt.Printf("fetch active users:")
			ListActiveUsers("LandLord.json")
		case "3":
			fmt.Printf("fetch Inactive users:")
			ListInactiveUsers("LandLord.json")
		case "4":
			username := utils.ReadInput("Enter username to delete: ")
			DeleteUser(username, "LandLord.json")
		case "5":
			fmt.Println("Exiting Admin Dashboard...")
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
