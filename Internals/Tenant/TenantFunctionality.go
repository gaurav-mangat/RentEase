package Tenant

import (
	"RentEase/Internals/Landlord"
	"RentEase/utils"
	"fmt"
	"strings"
)

func searchProperties() {
	var propertyType, pincode int
	var state, city, locality, choice string

	fmt.Println("Enter the type of property to search (1. Commercial, 2. Flat, 3. House):")
	fmt.Scan(&propertyType)

	state = utils.ReadInput("Enter the State (Full name):")
	city = utils.ReadInput("Enter the city:")
	locality = utils.ReadInput("Enter the area:")
	pincode = utils.ReadPincode()
	choice = utils.ReadInput("Enter 's' to search or 'f' to apply filters:")

	if choice == "s" {
		SearchPropertiesinFile(propertyType, state, city, locality, pincode)
	} else if choice == "f" {
		ApplyFilters(propertyType, state, city, locality, pincode)
	} else {
		fmt.Println("Please enter a valid choice.")
	}
}

func ApplyFilters(propertyType int, state, city, locality string, pincode int) {
	var minpriceRange, maxpriceRange float64
	var bhk string

	if propertyType == 3 {
		bhk = utils.ReadInput("Enter the BHK:") + " BHK"
	}

	fmt.Println("Enter the price range:")
	fmt.Println("Minimum value:")
	fmt.Scanf("%f", &minpriceRange)
	fmt.Println("Maximum value:")
	fmt.Scanf("%f", &maxpriceRange)

	// Apply filters and call the search function
	SearchPropertiesinFileWithFilters(propertyType, state, city, locality, pincode, bhk, minpriceRange, maxpriceRange)
}

func SearchPropertiesinFile(propertyType int, state, city, locality string, pincode int) {
	// Load properties from file
	_, err := Landlord.LoadProperties("prince.json")
	if err != nil {
		fmt.Println("Error loading properties:", err)
		return
	}

	// Search logic
	for _, property := range Landlord.Properties {
		if property.PropertyType == propertyType &&
			strings.Contains(property.Address.City, city) && property.Address.Area == locality && property.Address.Pincode == pincode {
			fmt.Printf("Property ID: %d, Title: %s, Price: %f\n", property.PropertyID, property.Title, property.Price)
		}
	}
}

func SearchPropertiesinFileWithFilters(propertyType int, state, city, locality string, pincode int, bhk string, minpriceRange, maxpriceRange float64) {
	// Load properties from file
	properties, err := Landlord.LoadProperties("prince.json")
	if err != nil {
		fmt.Println("Error loading properties:", err)
		return
	}

	// Filtered search logic
	for _, property := range properties {
		if property.PropertyType == propertyType &&
			property.Address.City == city && property.Address.Area == locality && property.Address.Pincode == pincode &&
			property.Price >= minpriceRange && property.Price <= maxpriceRange {
			fmt.Printf("Property ID: %d, Title: %s, Price: %f\n", property.PropertyID, property.Title, property.Price)
		}
	}
}
