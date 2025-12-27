package service

import (
	"auth-service/apihelpers"
	"auth-service/constants"
	"auth-service/db"
	"auth-service/dbops"
	"auth-service/kafka"
	"auth-service/loggerconfig"
	"auth-service/models"
	"context"
	"crypto/rand"
	"math/big"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func Signup(payload models.SignUp) (int, apihelpers.APIRes) {
	var apiRes apihelpers.APIRes

	userProfile := models.UserProfile{
		UserId:      uuid.New().String(),
		UserName:    payload.EmailId,
		Name:        payload.Name,
		Email:       payload.EmailId,
		PhoneNumber: payload.PhoneNumber,
		Password:    payload.Password,
		Verified:    false,
	}

	err := db.CreateUserProfile(userProfile)
	if err != nil {
		return apihelpers.SendInternalServerError("Some internal server occurred!")
	}
	return http.StatusOK, apiRes
}

func SendEmailOTP(email string) (int, apihelpers.APIRes) {
	var apiRes apihelpers.APIRes
	ctx := context.Background()

	otp, err := generateOTP(constants.OtpLengthForVerifyEmail)
	if err != nil {
		otp = "789451"
	}

	err = dbops.RedisSetValue(ctx, email, constants.OtpField, []byte(otp), 2*time.Minute)
	if err != nil {
		return apihelpers.SendInternalServerError("Redis error in setting otp-val with error: " + err.Error())
	}

	content := "Your otp is : " + otp
	loggerconfig.Info("OTP for email: ", email, " is ", otp)

	err = sendEmail(email, content)
	if err != nil {
		return apihelpers.SendInternalServerError("")
	}

	apiRes.Status = true
	apiRes.Message = constants.Success
	return http.StatusOK, apiRes
}

func VerifyEmailOtp(email, otp string) (int, apihelpers.APIRes) {
	var apiRes apihelpers.APIRes
	ctx := context.Background()

	orignalOtp, err := dbops.RedisGetValue(ctx, email, constants.OtpField)
	if err != redis.Nil && err != nil {
		loggerconfig.Info("VerifyEmailOtp (service) failed to fetched otp from redis with error: ", err)
		return apihelpers.SendInternalServerError("Failed to fetched otp from redis with error: " + err.Error())
	} else if err == redis.Nil {
		loggerconfig.Info("VerifyEmailOtp (service) no otp found in redis for email :", email)
		apiRes.Status = true
		apiRes.Message = "No otp found in redis for email: " + email
		return http.StatusBadRequest, apiRes
	}

	if orignalOtp != otp {
		loggerconfig.Info("Otp mismatched orginal otp :", orignalOtp, " otp got :", otp, " for email: ", email)
		apiRes.Status = true
		apiRes.Message = "otp mismatched"
		return http.StatusBadRequest, apiRes
	}

	err = sendNotification("", email, "SUCCESS")
	if err != nil {
		loggerconfig.Error("Failed to send notification with error: " + err.Error())
	}

	apiRes.Status = true
	apiRes.Message = constants.Success
	return http.StatusOK, apiRes
}

func generateOTP(digitCount int) (string, error) {
	const digits = "0123456789"
	length := digitCount
	b := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		b[i] = digits[randomIndex.Int64()]
	}
	return string(b), nil
}

func sendEmail(email, content string) error {
	producer := kafka.GetProducer()

	msg := models.EmailMessage{
		SenderEmail: email,
		Content:     content,
	}
	err := producer.SendEmail(context.Background(), msg)
	if err != nil {
		loggerconfig.Error("Failed to send email to kafka with error: ", err)
		return err
	}
	return nil
}

func sendNotification(userId, email, content string) error {
	producer := kafka.GetProducer()
	msg := models.NotificationMessage{
		UserId:  userId,
		Email:   email,
		Content: content,
	}

	err := producer.SendNotification(context.Background(), msg)
	if err != nil {
		loggerconfig.Error("Failed to send notification to kafka with error: ", err)
		return err
	}

	return nil
}
