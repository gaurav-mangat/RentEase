package utils

import (
	"RentEase/models"
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
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

// Global list of all property files

var filename = "Filenames.json"
var Filenames []string

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

// Reading pincode
func ReadPincode() int {
	var pincode int

	for {
		fmt.Print("Enter a 6-digit pincode: ")
		_, err := fmt.Scan(&pincode)

		// Check if there's an error in scanning or if the pincode is not 6 digits
		if err != nil || pincode < 100000 || pincode > 999999 {
			fmt.Println("Invalid pincode. Please enter a valid 6-digit pincode.")
			continue
		}

		// If valid, break the loop
		break
	}

	return pincode
}

// SaveAndUpdateFilenames saves the file path to the Filenames.json file and returns an updated global list of all filenames
func SaveAndUpdateFilenames(filePath string) ([]string, error) {
	// Load existing filenames
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// If the Filenames.json file doesn't exist, create it with an empty array
		err := ioutil.WriteFile(filename, []byte("[]"), 0644)
		if err != nil {
			return nil, fmt.Errorf("could not create Filenames.json: %v", err)
		}
	}

	// Read the existing filenames
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read Filenames.json: %v", err)
	}

	// Unmarshal the JSON data into the Filenames slice
	err = json.Unmarshal(data, &Filenames)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON data: %v", err)
	}

	// Check if the filename is already in the list to avoid duplicates
	for _, name := range Filenames {
		if name == filePath {
			// Return the existing list if the file path is already saved
			return Filenames, nil
		}
	}

	// Add the new file path to the list
	Filenames = append(Filenames, filePath)

	// Marshal the updated list back to JSON
	updatedData, err := json.Marshal(Filenames)
	if err != nil {
		return nil, fmt.Errorf("could not marshal JSON data: %v", err)
	}

	// Write the updated list back to Filenames.json
	err = ioutil.WriteFile(filename, updatedData, 0644)
	if err != nil {
		return nil, fmt.Errorf("could not write to Filenames.json: %v", err)
	}

	return Filenames, nil
}
