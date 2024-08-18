package models

// Abstract Class for User
type User struct {
	Username     string `json:"username"`
	PasswordHash string `json:"passwordHash"`
	Name         string `json:"name"`
	Age          int    `json:"age"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	Role         string `json:"role"`
	UserID       int    `json:"user_id"`
	UserStatus   int    `json:"user_status"`
}

// Struct for Landlord by embedding User struct
type Landlord struct {
	User
}

// Struct for Tenant by embedding User struct
type Tenant struct {
	User
}

// Struct for Admin by embedding User struct
type Admin struct {
	User
}

// Struct for Address

type Address struct {
	State   string `json:"state"`
	City    string `json:"city"`
	Area    string `json:"street"`
	Pincode int    `json:"pincode"`
}

// Struct for Property
type Property struct {
	PropertyID   int      `json:"propertyID"`
	PropertyType int      `json:"propertyType"`
	Title        string   `json:"title"`
	Address      Address  `json:"address"`
	LandlordID   int      `json:"landlordID"`
	Price        float64  `json:"price"`
	Availability int      `json:"availability"`
	Amenities    []string `json:"amenities"`
	RentalTerms  string   `json:"rentalTerms"`
}

// Struct for PropertyFilenames
type PropertyFilenames struct {
	Filename string `json:"filename"`
}

type Wishlist struct {
	Username   string // Assuming you track users by their username
	Properties []int  // List of PropertyIDs
}
type WishlistData map[string][]int // Map username to list of property IDs
