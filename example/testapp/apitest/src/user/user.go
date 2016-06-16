package user

// File with basic user structs

// a simple Users which has simple fields
type User struct {
	Name    string
	Age     int
	Friends []User
}
