package service

import (
	"auth-service/apihelpers"
	"auth-service/db"
	"auth-service/models"
	"net/http"

	"github.com/google/uuid"
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
