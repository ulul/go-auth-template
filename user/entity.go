package user

import "time"

// User object
type User struct {
	ID        int
	Name      string
	Email     string
	Ocupation string
	Password  string
	Avatar    string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
