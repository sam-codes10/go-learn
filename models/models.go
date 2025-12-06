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
