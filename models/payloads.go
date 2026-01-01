package models

type SignUp struct {
	Name        string `json:"name"`
	EmailId     string `json:"emailId"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type Login struct {
	EmailId  string `json:"emailId"`
	Password string `json:"password"`
}
