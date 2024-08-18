package InputHandlers

import (
	"RentEase/utils"
	"fmt"
	"regexp"
	"strconv"
)

func GetUsername() string {
	var username string
	for {
		username = utils.ReadInput("\033[1;34mEnter username (Username should only be a single word): \033[0m")
		if utils.IsValidInput(username) && utils.IsUsernameUnique(username) {
			return username
		}
		fmt.Println("\033[1;31mInvalid or duplicate username.\033[0m")
	}
}

func GetPassword() string {
	var password string
	for {
		password = utils.ReadInput("\033[1;34m\nEnter password (min 9 chars, include lowercase, uppercase, numbers, special): \033[0m")
		if utils.IsValidInput(password) && utils.IsValidPassword(password) {
			return password
		}
		fmt.Println("\033[1;31m\nPassword does not meet complexity requirements.\033[0m")
	}
}

func GetFullName() string {
	return utils.ReadInput("\033[1;34m\nEnter full name: \033[0m")
}

func GetMobileNumber() string {
	var mobileNumber string
	for {
		mobileNumber = utils.ReadInput("\033[1;34m\nEnter mobile number: \033[0m")
		if utils.IsValidInput(mobileNumber) && utils.IsValidMobileNumber(mobileNumber) {
			return mobileNumber
		}
		fmt.Println("\033[1;31m\nInvalid mobile number.\033[0m")
	}
}

func GetAddress() string {
	return utils.ReadInput("\033[1;34m\nEnter address: \033[0m")
}

func GenerateUniqueUserID() int {
	// Example logic: generate unique userID
	return len(utils.Users) + 1
}

func GetAge() int {
	for {

		ageStr := utils.ReadInput("Enter Age: ") // Assumes readInput() reads a string from standard input

		// Try to convert the input to an integer
		age, err := strconv.Atoi(ageStr)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		// Validate that the age is within a reasonable range
		if age > 21 && age <= 120 {
			return age
		}

		fmt.Println("Invalid age. Please enter a valid age between 19 and 120.")
	}
}
func GetEmail() string {
	for {
		email := utils.ReadInput("Enter Email: ") // Assumes readInput() reads a string from standard input

		// Validate the email format
		if isValidEmail(email) {
			return email
		}

		fmt.Println("Invalid email address. Please enter a valid email address.")
	}
}

// isValidEmail validates the email format using a regular expression
func isValidEmail(email string) bool {
	// Basic email validation regex
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
