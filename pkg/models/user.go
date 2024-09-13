package models

import "time"

type User struct {
	ID        int
	UserName  string
	Password  string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
	Email     string
	Phone     string
}
