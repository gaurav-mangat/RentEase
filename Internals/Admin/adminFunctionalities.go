package Admin

import (
	"RentEase/models"
	"encoding/json"
	"fmt"
	"os"
)

// LoadUsers loads all users from the JSON file
func LoadUsers(filePath string) ([]models.User, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var users []models.User
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&users)
	if err != nil {
		return nil, fmt.Errorf("could not decode JSON data: %v", err)
	}
	//for _, user := range users {
	//	if user.UserStatus == 0 || user.UserStatus == 1 { // Active or Inactive user
	//		fmt.Printf("Username: %s, Full Name: %s, UserID: %d\n", user.Username, user.Name, user.UserID)
	//	}
	//}

	return users, nil
}

// ListActiveUsers lists all active users  0 - active 1 - inactive
func ListActiveUsers(filePath string) {
	users, err := LoadUsers(filePath)
	if err != nil {
		fmt.Println("Error loading users:", err)
		return
	}

	fmt.Println("Active Users:")
	for _, user := range users {
		if user.UserStatus == 0 { // Active
			fmt.Printf("Username: %s, Full Name: %s, UserID: %d\n", user.Username, user.Name, user.UserID)
		}
	}
}

func LoadAllUsers() {
	landlords, err := LoadUsers("LandLord.json")
	if err != nil {
		fmt.Println("Error loading landlords:", err)
		return
	}

	tenants, err := LoadUsers("Tenant.json")
	if err != nil {
		fmt.Println("Error loading tenants:", err)
		return
	}

	fmt.Println("Landlords:")
	for _, landlord := range landlords {
		fmt.Printf("UserID: %d, Username: %s, Name: %s\n", landlord.UserID, landlord.Username, landlord.Name)
	}

	fmt.Println("Tenants:")
	for _, tenant := range tenants {
		fmt.Printf("UserID: %d, Username: %s, Name: %s\n", tenant.UserID, tenant.Username, tenant.Name)
	}
}

// ListInactiveUsers lists all inactive users
func ListInactiveUsers(filePath string) {
	users, err := LoadUsers(filePath)
	if err != nil {
		fmt.Println("Error loading users:", err)
		return
	}

	fmt.Println("Inactive Users:")
	for _, user := range users {
		if user.UserStatus == 1 { // Inactive
			fmt.Printf("Username: %s, Full Name: %s, UserID: %d\n", user.Username, user.Name, user.UserID)
		}
	}
}

// DeleteUser deletes a user by setting their status to inactive
func DeleteUser(username string, filePath string) {
	users, err := LoadUsers(filePath)
	if err != nil {
		fmt.Println("Error loading users:", err)
		return
	}

	var userFound bool
	for i, user := range users {
		if user.Username == username {
			users[i].UserStatus = 1 // Mark user as inactive
			userFound = true
			break
		}
	}

	if !userFound {
		fmt.Println("User not found.")
		return
	}

	// Save updated users list
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(users)
	if err != nil {
		fmt.Println("Error saving updated users:", err)
		return
	}

	fmt.Println("User deleted successfully!")
}
