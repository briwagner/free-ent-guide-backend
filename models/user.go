package models

// User model.
type User struct {
	Name string   `json:"username,omitempty"`
	Zips []string `json:"zipcodes"`
}
