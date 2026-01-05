package db

import (
	"auth-service/constants"
	"auth-service/dbops"
	"auth-service/models"
	"context"
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

func CheckEmailIsAlreadyInUseByUser(email string, ctx context.Context) (bool, error) {
	count := 0

	query := `SELECT COUNT(*) FROM user_profile WHERE email = ? AND role = ?`

	err := dbops.DB.WithContext(ctx).Raw(query, email, constants.RoleUser).Scan(&count).Error
	if err != nil {
		return false, nil
	}

	if count == 0 {
		return false, nil
	}
	return true, nil
}

func GetUsernameCount(username string, ctx context.Context) (int, error) {
	count := 0

	query := `SELECT COUNT(*) FROM user_profile WHERE user_name = ? AND role = ?`

	err := dbops.DB.WithContext(ctx).Raw(query, username, constants.RoleGuest).Scan(&count).Error
	if err != nil {
		return count, err
	}

	if count == 0 {
		return count, err
	}

	return count, nil
}
