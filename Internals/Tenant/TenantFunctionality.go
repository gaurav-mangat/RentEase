package Tenant

import (
	"RentEase/Internals/Landlord"
	"RentEase/models"
	"RentEase/utils"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func searchProperties() {
	var propertyType, pincode int
	var state, city, locality, choice string

	fmt.Println("Enter the type of property to search (1. Commercial, 2. Flat, 3. House):")
	fmt.Scan(&propertyType)

	state = strings.TrimSpace(strings.ToLower(utils.ReadInput("Enter the State (Full name):")))
	city = strings.TrimSpace(strings.ToLower(utils.ReadInput("Enter the city:")))
	locality = strings.TrimSpace(strings.ToLower(utils.ReadInput("Enter the area:")))
	pincode = utils.ReadPincode()
	choice = strings.TrimSpace(strings.ToLower(utils.ReadInput("Enter 's' to search or 'f' to apply filters:")))

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

	if propertyType == 2 {
		bhk = strings.TrimSpace(utils.ReadInput("Enter the BHK:")) + " BHK"
	}

	fmt.Println("Enter the price range:")
	fmt.Println("Minimum value:")
	fmt.Scanf("%f", &minpriceRange)
	fmt.Println("Maximum value:")
	fmt.Scanf("%f", &maxpriceRange)

	// Apply filters and call the search function
	propertyIDs := SearchPropertiesinFileWithFilters(propertyType, state, city, locality, pincode, bhk, minpriceRange, maxpriceRange, 0)
	fmt.Println(propertyIDs)
}

func SearchPropertiesinFile(propertyType int, state, city, locality string, pincode int) {
	// Load properties from file
	properties, err := Landlord.LoadProperties("prince.json")
	if err != nil {
		fmt.Println("Error loading properties:", err)
		return
	}

	// Search logic
	found := false
	for _, property := range properties {
		if property.PropertyType == propertyType &&
			strings.Contains(strings.ToLower(property.Address.City), city) &&
			property.Address.Pincode == pincode {
			fmt.Println("\n Searched Properties :")
			fmt.Printf("Property ID: %d, Title: %s, Price: %f\n", property.PropertyID, property.Title, property.Price)
			found = true
		}
	}
	if !found {
		fmt.Println("No property found!")
	}
}

func SearchPropertiesinFileWithFilters(propertyType int, state, city, locality string, pincode int, bhk string, minpriceRange, maxpriceRange float64, propertyid int) []int {
	// Load properties from file
	properties, err := Landlord.LoadProperties("prince.json")
	if err != nil {
		fmt.Println("Error loading properties:", err)
		return nil
	}

	var listOfProperties []int

	// Filtered search logic
	for _, property := range properties {
		// Filtered search logic
		if (property.PropertyType == propertyType &&
			strings.ToLower(property.Address.City) == city &&
			property.Address.Pincode == pincode &&
			property.Price >= minpriceRange && property.Price <= maxpriceRange) || (property.PropertyID == propertyid && property.PropertyID != 0) {
			fmt.Println("\n Searched Properties :")
			fmt.Printf("Property ID: %d, Title: %s, Price: %f\n", property.PropertyID, property.Title, property.Price)
		}
		matches := true
		if propertyType != 0 && property.PropertyType != propertyType {
			matches = false
		}
		if state != "" && strings.ToLower(property.Address.State) != state {
			matches = false
		}
		if city != "" && strings.ToLower(property.Address.City) != city {
			matches = false
		}
		if locality != "" && strings.ToLower(property.Address.Area) != locality {
			matches = false
		}
		if pincode != 0 && property.Address.Pincode != pincode {
			matches = false
		}
		if minpriceRange != 0 && maxpriceRange != 0 && (property.Price < minpriceRange || property.Price > maxpriceRange) {
			matches = false
		}
		if propertyid != 0 && property.PropertyID != propertyid {
			matches = false
		}

		if matches {
			listOfProperties = append(listOfProperties, property.PropertyID)
		}
	}

	return listOfProperties
}

// File where the wishlist data is stored
const wishlistFile = "wishlist.json"

func LoadWishlistData() (models.WishlistData, error) {
	file, err := os.Open(wishlistFile)
	if err != nil {
		if os.IsNotExist(err) {
			return models.WishlistData{}, nil // File doesn't exist, return empty map
		}
		return nil, err
	}
	defer file.Close()

	var data models.WishlistData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func SaveWishlistData(data models.WishlistData) error {
	file, err := os.Create(wishlistFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // For pretty-printing
	return encoder.Encode(data)
}

// Add to wishlist
func AddToWishlist(propertyID int, username string) {
	propertyIDs := SearchPropertiesinFileWithFilters(0, "", "", "", 0, "", 0, 0, propertyID)

	if !Contains(propertyIDs, propertyID) {
		fmt.Println("Property ID not found in the search results.")
		return
	}

	data, err := LoadWishlistData()
	if err != nil {
		fmt.Println("Error loading wishlist data:", err)
		return
	}

	// Add property ID to the user's wishlist
	if data[username] == nil {
		data[username] = []int{}
	}
	for _, id := range data[username] {
		if id == propertyID {
			fmt.Println("Property is already in your wishlist.")
			return
		}
	}
	data[username] = append(data[username], propertyID)

	err = SaveWishlistData(data)
	if err != nil {
		fmt.Println("Error saving wishlist data:", err)
		return
	}

	fmt.Println("Property added to your wishlist!")
}

// View wishlist
func ViewWishlist(username string) {
	data, err := LoadWishlistData()
	if err != nil {
		fmt.Println("Error loading wishlist data:", err)
		return
	}

	propertyIDs := data[username]
	if propertyIDs == nil {
		fmt.Println("Your wishlist is empty.")
		return
	}

	// Load properties from file
	properties, err := Landlord.LoadProperties("prince.json")
	if err != nil {
		fmt.Println("Error loading properties:", err)
		return
	}

	fmt.Println("Your Wishlist:")
	for _, propertyID := range propertyIDs {
		for _, property := range properties {
			if property.PropertyID == propertyID {
				fmt.Printf("Property ID: %d, Title: %s, Price: %f, City: %s, Area: %s, Pincode: %d\n",
					property.PropertyID, property.Title, property.Price, property.Address.City,
					property.Address.Area, property.Address.Pincode)
				break
			}
		}
	}
}

func Contains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
