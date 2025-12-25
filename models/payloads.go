package models

type SignUp struct {
	Name        string `json:"name"`
	EmailId     string `json:"emailId"`
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type ConsumerMessageOTP struct {
	Email           bool   `json:"email"`
	SMS             bool   `json:"sms"`
	Content         string `json:"content"`
	Subject         string `json:"subject"`
	RecieverMail    string `json:"recieverMail"`
	RecieverContact string `json:"recieverContact"`
}
