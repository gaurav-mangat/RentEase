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
	Area    string `json:"area"`
	Street  string `json:"street"`
	Pincode string `json:"pincode"`
}

// Struct for Property
type Property struct {
	PropertyID   int      `json:"propertyID"`
	PropertyType string   `json:"propertyType"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Address      Address  `json:"address"`
	LandlordID   int      `json:"landlordID"`
	Price        float64  `json:"price"`
	Availability string   `json:"availability"`
	Amenities    []string `json:"amenities"`
	RentalTerms  string   `json:"rentalTerms"`
}
