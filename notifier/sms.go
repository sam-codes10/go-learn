package notifier

import (
	"encoding/json"
	"fmt"
	"log"
	"notifier-service/models"
)

func SendSMS(value []byte) error {
	var msg models.SMSMessage

	err := json.Unmarshal(value, &msg)
	if err != nil {
		log.Fatal("SendSMS (notifier) - Wrong SMS format")
		return err
	}
	fmt.Println("sending email to:", msg.To)
	return nil
}
