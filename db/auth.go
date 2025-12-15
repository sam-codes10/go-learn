package db

import (
	"auth-service/dbops"
	"auth-service/models"
)

func CreateUserProfile(userProfile models.UserProfile) error {
	return dbops.DB.Create(userProfile).Error
}
