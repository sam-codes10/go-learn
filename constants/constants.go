package constants

const OtpLengthForVerifyEmail = 6

const (
	OtpField = "otp-field"
)

const (
	Success = "SUCCESS"
	Fail    = "FAIL"
)

// kafka-topic names
const (
	EmailTopic        = "email"
	NotificationTopic = "notification"
)

var DomainEmailMapping = map[string]string{
	"gmail.com":   "g",
	"yahoo.com":   "y",
	"hotmail.com": "h",
}

const (
	RoleGuest = "guest"
	RoleUser  = "user"
	RoleAdmin = "admin"
)
