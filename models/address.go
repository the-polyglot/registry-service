package models

type Address struct {
	addressLine string
	city        string
	street      string
	postalCode  string
	email       string
}

func NewAddress(addressLine, city, street, postalCode, email string) Address {
	return Address{addressLine, city, street, postalCode, email}
}
