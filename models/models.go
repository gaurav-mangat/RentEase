package models

// Abstract Class for User
type User struct {
	username    string
	password    string
	name        string
	age         int
	email       string
	phoneNumber string
	address     string
	role        string
	userID      int
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
