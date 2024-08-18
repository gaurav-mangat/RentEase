package utils

import (
	"RentEase/models"
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Create a buffered reader
var Reader *bufio.Reader

// Initialize the Reader in the init function
func init() {
	Reader = bufio.NewReader(os.Stdin)
}

// Global variable for user storage
var Users []models.User

// Global variable for property storage and ID tracking

var NextPropertyID int

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

// LoadUsers loads users from a JSON file, creating the file if it doesn't exist.
func LoadUsers(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File does not exist, create an empty file
			err = os.WriteFile(filename, []byte("[]"), 0644)
			if err != nil {
				return fmt.Errorf("failed to create file: %v", err)
			}
			Users = []models.User{}
			return nil
		}
		return err
	}

	// Check if the file is empty
	if len(file) == 0 {
		Users = []models.User{}
		return nil
	}

	// Try to unmarshal the file content into users slice
	if err := json.Unmarshal(file, &Users); err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		Users = []models.User{} // Reset Users to empty slice
		return err
	}
	return nil
}

func SaveUsers(filename string) error {
	file, err := json.MarshalIndent(Users, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}

// ReadFloat reads a float value from the user.
func ReadFloat(prompt string) float64 {
	for {
		input := ReadInput(prompt)
		value, err := strconv.ParseFloat(input, 64)
		if err == nil {
			return value
		}
		fmt.Println("Invalid input. Please enter a valid float value.")
	}
}

// ParseCommaSeparatedList parses a comma-separated list of strings into a slice.
func ParseCommaSeparatedList(input string) []string {
	items := strings.Split(input, ",")
	for i := range items {
		items[i] = strings.TrimSpace(items[i])
	}
	return items
}
