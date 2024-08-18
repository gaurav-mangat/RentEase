package Authentication

import (
	"RentEase/Internals/InputHandlers"
	"RentEase/models"
	"RentEase/utils"
	"fmt"
)

// Sign up a new user

func SignUp(userType string) {
	// Load users from the file based on the type of user
	var file string
	if userType == "Landlord" {
		file = "LandLord.json"
	} else if userType == "Tenant" {
		file = "Tenant.json"
	}

	if err := utils.LoadUsers(file); err != nil {
		fmt.Printf("\033[1;31mError loading users: %v\033[0m\n", err) // Red bold
		return
	}

	// Collect user details

	username := InputHandlers.GetUsername()
	password := InputHandlers.GetPassword()
	fullName := InputHandlers.GetFullName()
	age := InputHandlers.GetAge()
	mobileNumber := InputHandlers.GetMobileNumber()
	email := InputHandlers.GetEmail()
	address := InputHandlers.GetAddress()

	// Hash password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Printf("\033[1;31m\nError hashing password: %v\033[0m\n", err) // Red bold
		return
	}

	// Create User struct
	userID := InputHandlers.GenerateUniqueUserID()
	newUser := models.User{
		Username:     username,
		PasswordHash: hashedPassword,
		Name:         fullName,
		Age:          age,
		Email:        email,
		PhoneNumber:  mobileNumber,
		Address:      address,
		Role:         userType,
		UserID:       userID,
	}

	// Adding a new user to slice of user
	utils.Users = append(utils.Users, newUser)

	// Save user
	if err := utils.SaveUsers(file); err != nil {
		fmt.Printf("\033[1;31mError saving user: %v\033[0m\n", err) // Red bold
	} else {
		fmt.Println("\033[1;32m\n\nUser signed up successfully!\033[0m") // Green bold
		promptPostSignUp(userType)
	}
}

func promptPostSignUp(userType string) {
	fmt.Println("\n\nPress 1 to Login \nPress 2 to Exit")
	var choice int
	fmt.Print("\033[1;34m\nEnter your choice: \033[0m") // Blue bold
	_, err := fmt.Scan(&choice)
	if err != nil {
		fmt.Printf("\033[1;31mError reading choice: %v\033[0m\n", err) // Red bold
		return
	}

	switch choice {
	case 1:
		Login(userType)
	case 2:
		return
	default:
		fmt.Println("\033[1;31mInvalid choice\033[0m") // Red bold
	}
}
