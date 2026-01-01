package models

type Claims struct {
	Uuid       string
	Email      string
	Role       string
	IsVerified bool
}

const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)
