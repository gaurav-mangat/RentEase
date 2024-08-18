package Landlord

import (
	"RentEase/models"
	"RentEase/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Global active Landlord
var activeLandlord models.User

var Properties []models.Property

// Global File Path
var Filepath string

// LoadProperties function loads properties from a JSON file, creates the file if it doesn't exist
func LoadProperties(filePath string) ([]models.Property, error) {
	Filepath = filePath
	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// If the file doesn't exist, create it with an empty JSON array
		err := ioutil.WriteFile(filePath, []byte("[]"), 0644)
		if err != nil {
			return nil, fmt.Errorf("could not create file: %v", err)
		}
	}

	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	// Read the file contents
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	// Unmarshal the JSON data into the properties slice
	err = json.Unmarshal(data, &Properties)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON data: %v", err)
	}

	// Saving the file name to it
	utils.Filenames, err = utils.SaveAndUpdateFilenames(Filepath)
	if err != nil {
		return nil, fmt.Errorf("could not save filename: %v", err)
	}
	return Properties, nil
}

// SaveProperties saves the property list to a JSON file
func SaveProperties(filePath string, properties []models.Property) error {
	file, err := json.MarshalIndent(properties, "", "    ")
	if err != nil {
		return fmt.Errorf("could not marshal properties: %v", err)
	}

	err = os.WriteFile(filePath, file, 0644)
	if err != nil {
		return fmt.Errorf("could not write file: %v", err)
	}

	return nil
}

// addProperties adds a new property to the user's property list and saves it to a JSON file
func addProperties() {

	// Collect new property details from the user
	fmt.Println("Enter details for the new property:")

	var propertyType int
	fmt.Println("Property type :\n1. Commercial\n2.Flats\n3.House")
	var ch int
	fmt.Scan(&ch)
	switch ch {
	case 1:
		propertyType = 1
	case 2:
		propertyType = 2
	case 3:
		propertyType = 3
	}
	title := utils.ReadInput("Title: ")
	area := utils.ReadInput("Area: ")
	city := utils.ReadInput("City: ")
	state := utils.ReadInput("State: ")
	pincode := utils.ReadPincode()
	price := utils.ReadFloat("Price: ")
	availability := 0
	amenities := utils.ParseCommaSeparatedList(utils.ReadInput("Amenities (comma-separated): "))
	rentalTerms := utils.ReadInput("Rental Terms: ")

	// Create a new property
	newProperty := models.Property{
		PropertyID: GenerateUniquePropertyID(),
		Title:      title,
		Address: models.Address{
			Area:    area,
			City:    city,
			State:   state,
			Pincode: pincode,
		},
		LandlordID:   activeLandlord.UserID,
		Price:        price,
		Availability: availability,
		PropertyType: propertyType,
		Amenities:    amenities,
		RentalTerms:  rentalTerms,
	}

	// Add the new property to the list and save it
	Properties = append(Properties, newProperty)
	err := SaveProperties(Filepath, Properties)
	if err != nil {
		fmt.Println("Error saving property:", err)
	} else {
		fmt.Println("Property added successfully!")
	}
}

func GenerateUniquePropertyID() int {
	// Example logic: generate unique userID
	return len(Properties) + 1
}

func deleteProperty() {

	if len(Properties) == 0 {
		fmt.Println("You have no properties to delete.")
		return
	}

	fmt.Println("Available properties:")
	for _, property := range Properties {
		fmt.Printf("ID: %d - Title: %s\n", property.PropertyID, property.Title)
	}

	fmt.Print("Enter the ID of the property you want to delete: ")
	var propertyID int
	_, err := fmt.Scan(&propertyID)
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}

	// Find and delete the property with the given ID
	var propertyFound bool
	for i, property := range Properties {
		if property.PropertyID == propertyID {
			// Delete the property by slicing out the element
			Properties = append(Properties[:i], Properties[i+1:]...)
			propertyFound = true
			break
		}
	}

	if !propertyFound {
		fmt.Println("Property not found.")
		return
	}

	// Save the updated properties back to the file
	err = SaveProperties(Filepath, Properties)
	if err != nil {
		fmt.Println("Error saving property:", err)
	} else {
		fmt.Println("Property deleted successfully!")
	}
}

func viewProfile() {
	fmt.Println()
	fmt.Println()
	fmt.Println("\033[1;36m---------------------------------------------\033[0m") // Sky blue
	fmt.Println("\033[1;34m               YOUR PROFILE                 \033[0m")  // Blue
	fmt.Println("\033[1;36m---------------------------------------------\033[0m")
	fmt.Printf("Username       : %s\n", activeLandlord.Username)
	fmt.Printf("Full Name      : %s\n", activeLandlord.Name)
	fmt.Printf("Age            : %s\n", activeLandlord.Age)
	fmt.Printf("Email          : %s\n", activeLandlord.Email)
	fmt.Printf("Phone Number   : %s\n", activeLandlord.PhoneNumber)
	fmt.Printf("Address        : %s\n", activeLandlord.Address)
	fmt.Println("\033[1;36m---------------------------------------------\033[0m")

	fmt.Println("Your listed properties are :")
	for i, property := range Properties {
		fmt.Println("\nProperty Number ", i+1, " :")
		fmt.Printf("ID: %d\n", property.PropertyID)
		pT := property.PropertyType
		switch pT {
		case 1:
			fmt.Println("Property Type: Commercial ")

		case 2:
			fmt.Println("Property Type: Flats ")

		case 3:
			fmt.Println("Property Type: House ")
		}
		fmt.Printf("Title: %s\n", property.Title)
		fmt.Printf("Address: %s, %s, %s, %s\n",
			property.Address.Area,
			property.Address.City,
			property.Address.State,
			property.Address.Pincode)
		fmt.Printf("Landlord ID: %d\n", property.LandlordID)
		fmt.Printf("Price: %.2f\n", property.Price)
		fmt.Printf("Availability: %d\n", property.Availability)
		fmt.Printf("Amenities: %v\n", property.Amenities)
		fmt.Printf("Rental Terms: %s\n", property.RentalTerms)

	}
	fmt.Println()
	fmt.Println()
}
