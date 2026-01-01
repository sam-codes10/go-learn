package db

import (
	"auth-service/dbops"
	"auth-service/models"
)

func CreateUserProfile(userProfile models.UserProfile) error {
	return dbops.DB.Create(userProfile).Error
}

func MarkUserEmailAsVerified(userId string) error {
	return dbops.DB.Model(&models.UserProfile{}).Where("user_id = ?", userId).Update("verified", true).Error
}

func GetUserProfileByEmail(email string) (models.UserProfile, error) {
	var userProfile models.UserProfile
	err := dbops.DB.Where("email_id = ?", email).First(&userProfile).Error
	return userProfile, err
}
