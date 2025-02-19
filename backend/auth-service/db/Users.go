package db

import "time"

type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	CreatedAt    time.Time `json:"createdAt"`
	LastLoggedIn time.Time `json:"lastLoggedIn"`
}
