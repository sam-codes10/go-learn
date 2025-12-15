package notifier

import (
	"encoding/json"
	"fmt"
	"log"
	"notifier-service/constants"
	"notifier-service/models"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendSMS(value []byte) error {
	var msg models.SMSMessage

	msgFrom := constants.MsgFrom

	err := json.Unmarshal(value, &msg)
	if err != nil {
		log.Fatal("SendSMS (notifier) - Wrong SMS format")
		return err
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: constants.TwilioAccountSID,
		Password: constants.TwilioAuthToken,
	})

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
