package models

type SignUp struct {
	Name        string `json:"name"`
	EmailId     string `json:"emailId"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}
