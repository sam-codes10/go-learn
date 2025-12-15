package db

import (
	"auth-service/dbops"
	"auth-service/models"
)

func SaveUserProfile(userProfile models.UserProfile) error {
	return dbops.DB.Save(userProfile).Error
}
