package models

type SignUp struct {
	Name        string `json:"name" validate:"required,min=2"`
	EmailId     string `json:"emailId" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Password    string `json:"password" validate:"required,min=8"`
}

type Login struct {
	EmailId  string `json:"emailId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role validate:"required,oneof=admin user guest"`
}
