package models

type Claims struct {
	Uuid       string
	Email      string
	Role       string
	IsVerified bool
}

const (
	UserRole  = "user"
	AdminRole = "admin"
)
