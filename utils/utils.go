package utils

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

var Reader bufio.Reader

// IsUsernameUnique checks if the username is unique.
func IsUsernameUnique(username string) bool {
	for _, user := range Users {
		if user.Username == username {

			fmt.Println("This username is already taken.")
			return false
		}
	}
	return true
}

func IsValidInput(input string) bool {
	if strings.Contains(input, " ") {

		return false
	}
	return true
}

func IsValidInput2(input string) bool {
	if strings.Contains(input, " ") {
		fmt.Println("\033[1;31m\nInvalid Input\033[0m")
		fmt.Println("\nTry again....")
		return false
	}
	return true
}

// IsValidPassword validates the password against specified criteria.
func IsValidPassword(password string) bool {
	var (
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString
		hasSpecial = regexp.MustCompile(`[!@#\$%\^&\*\(\)_+\-=\[\]\;:'",.<>?/|\\]`).MatchString
	)

	return len(password) > 8 && hasUpper(password) && hasLower(password) && hasNumber(password) && hasSpecial(password)
}

// HashPassword hashes a password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPasswordHash compares a hashed password with a plaintext password.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsValidMobileNumber(number string) bool {
	match, _ := regexp.MatchString(`^[6-9]\d{9}$`, number)
	return match
}

// ReadInput reads input from the user with a prompt.
func ReadInput(prompt string) string {
	fmt.Print(prompt)
	input, err := Reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return strings.TrimSpace(input)
}
