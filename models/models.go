package models

type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type ConsumerMessage struct {
	Content      string `json:"content"`
	Subject      string `json:"subject"`
	RecieverMail string `json:"recieverMail"`
}
