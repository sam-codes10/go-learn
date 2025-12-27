package models

type EmailMessage struct {
	SenderEmail string `json:"senderEmail"`
	Content     string `json:"content"`
}

type NotificationMessage struct {
	UserId  string `json:"userId"`
	Email   string `json:"email"`
	Content string `json:"content"`
}
