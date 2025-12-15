package notifier

import (
	"encoding/json"
	"fmt"
	"log"
	"notifier-service/models"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(value []byte) error {
	var msg models.SMSMessage
	msgFrom := os.Getenv("MSG_FROM")
	err := json.Unmarshal(value, &msg)
	if err != nil {
		log.Fatal("SendSMS (notifier) - Wrong SMS format")
		return err
	}

	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo(msg.To)
	params.SetFrom(msgFrom)
	params.SetBody(msg.Text)

	_, err = client.Api.CreateMessage(
		params,
	)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}
