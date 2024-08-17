package Landlord

import (
	"FileHandling/Config"
	"RentEase/utils"
	"fmt"
)

func Login() {
	const filename = Config.UserFile

	// Load users from the file once

	if err := utils.LoadUsers(Config.UserFile); err != nil {
		fmt.Printf("Error loading users: %v\n", err)
		return
	}

	attemptsLeft := 3

	for attemptsLeft > 0 {
		var username, password string
		fmt.Println()
		fmt.Println()
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m") // Sky blue
		fmt.Println("\033[1;31m                          LOG IN                                \033[0m") // Red bold
		fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")

		username = utils.ReadInput("\n             Enter username: ")
		if !utils.IsValidInput2(username) {
			return
		}

		password = utils.ReadInput("             Enter password: ")

		if !utils.IsValidInput(password) {
			return
		}
		fmt.Println()

		// Check credentials

		loginSuccessful := false
		for _, user := range utils.Users {
			if user.Username == username && utils.CheckPasswordHash(password, user.PasswordHash) {
				fmt.Println("\033[1;31m              Login successful!\033[0m") // Red bold
				loginSuccessful = true
				ol.ActiveUser = user
				ol.Dashboard(user)
				fmt.Println()
				return // Exit after successful login
			}
		}

		if loginSuccessful {
			// Successful login, exit the loop
			break
		} else {
			// Failed login, decrement the attempts left
			attemptsLeft--
			if attemptsLeft == 0 {
				fmt.Println("\033[1;31mLogin failed. You have exhausted all attempts.\033[0m") // Red bold
				return
			}
			fmt.Printf("Login failed. Please check your username and password. You have %d attempt(s) left.\n", attemptsLeft)

			fmt.Println("\nWhat would you like to do next?")
			fmt.Println("1. Retry Login")
			fmt.Println("2. Sign up")
			fmt.Println("3. Exit")

			var choice int
			fmt.Print("Enter your choice: ")
			fmt.Scan(&choice)

			switch choice {
			case 1:
				// Retry login
				continue
			case 2:
				// Call the SignUp function
				SignUp()
				return // Return to avoid retrying after sign up
			case 3:
				// Exit
				fmt.Println("Exiting...")
				return
			default:
				// Invalid choice
				fmt.Println("Invalid choice. Exiting...")
				return
			}
		}
	}
}

// Sign up a new user
func SignUp() {
	const filename = Config.UserFile

	// Load users from the file
	if err := utils.LoadUsers(Config.UserFile); err != nil {
		fmt.Printf("\033[1;31mError loading users: %v\033[0m\n", err) // Red bold
		return
	}

	// Signup form
	fmt.Println()
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")    // Sky blue
	fmt.Println("\033[1;31m                       SIGN UP FORM                                \033[0m") // Red bold
	fmt.Println("\033[1;36m----------------------------------------------------------------\033[0m")
	fmt.Println()

	// Get username
	var username string
	valid := false
	for !valid {
		username = utils.ReadInput("\033[1;34mEnter username (Username should only be a single word): \033[0m")
		if utils.IsValidInput(username) {
			valid = true
		} else {
			fmt.Println("\033[1;31mInvalid username.\nPlease enter a valid username.\n\033[0m")
		}
		if !utils.IsUsernameUnique(username) {
			return
		}

	}

	// Get password

	var password string
	valid = false
	for !valid {
		password = utils.ReadInput("\033[1;34m\nEnter password (min 9 chars, include lowercase, uppercase, numbers, special): \033[0m")
		if utils.IsValidInput(password) && utils.IsValidPassword(password) {
			valid = true
		} else {
			fmt.Println("\033[1;31m\nPassword does not meet complexity requirements.\nPlease enter a valid password.\033[0m")
		}
	}

	// Get full name

	var fullName string

	fullName = utils.ReadInput("\033[1;34m\nEnter full name: \033[0m")

	// Get mobile number
	var mobileNumber string
	valid = false
	for !valid {
		mobileNumber = utils.ReadInput("\033[1;34m\nEnter mobile number: \033[0m")
		if utils.IsValidInput(mobileNumber) && utils.IsValidMobileNumber(mobileNumber) {
			valid = true
		} else {
			fmt.Println("\033[1;31m\nInvalid mobile number.\nPlease enter a 10-digit number starting with 6, 7, 8, or 9.\033[0m")
		}
	}

	// Get gender
	var gender string
	valid = false
	for !valid {
		gender = utils.ReadInput("\033[1;34m\nEnter gender (Male/Female/Other): \033[0m")
		if utils.IsValidInput(gender) && (gender == "Male" || gender == "Female" || gender == "Others") {
			valid = true
		} else {
			fmt.Println("\033[1;31m\nInvalid gender.\nPlease enter a valid gender.\033[0m")
		}
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Printf("\033[1;31m\nError hashing password: %v\033[0m\n", err) // Red bold
		return
	}

	utils.Users = append(utils.Users, models.User{
		Username:     username,
		PasswordHash: hashedPassword,
		FullName:     fullName,
		MobileNumber: mobileNumber,
		Gender:       gender,
	})

	if err := utils.SaveUsers(Config.UserFile); err != nil {
		fmt.Printf("\033[1;31mError saving users: %v\033[0m\n", err) // Red bold
	} else {
		fmt.Println("\033[1;32m\n\nUser signed up successfully!\033[0m") // Green bold
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
			Login()
		case 2:
			return
		default:
			fmt.Println("\033[1;31mInvalid choice\033[0m") // Red bold
		}
	}
}
