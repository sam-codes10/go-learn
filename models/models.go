package models

type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type SMSMessage struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

type ConsumerMessageOTP struct {
	Email           bool   `json:"email"`
	SMS             bool   `json:"sms"`
	Content         string `json:"content"`
	Subject         string `json:"subject"`
	RecieverMail    string `json:"recieverMail"`
	RecieverContact string `json:"recieverContact"`
}
