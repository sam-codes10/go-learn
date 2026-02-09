package service

import (
	"auth-service/apihelpers"
	"auth-service/constants"
	"auth-service/db"
	"auth-service/dbops"
	"auth-service/helpers"
	"strconv"

	"auth-service/loggerconfig"
	"auth-service/middleware"
	"auth-service/models"
	"context"

	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

func Signup(payload models.SignUp) (int, apihelpers.APIRes) {
	ctx := context.Background()
	var apiRes apihelpers.APIRes

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		loggerconfig.Error("Signup (service) failed to hash password: ", err)
		return apihelpers.SendInternalServerError("Failed to hash password: " + err.Error())
	}

	userProfile := models.UserProfile{
		UserId:      uuid.New().String(),
		UserName:    helpers.ExtractUserNameFromEmail(payload.EmailId),
		Name:        payload.Name,
		Email:       payload.EmailId,
		PhoneNumber: payload.PhoneNumber,
		Password:    string(hashedPassword),
		// Verified:    false,
		Role:      constants.RoleGuest,
		Lock:      false,
		CreatedAt: time.Now().Unix(),
	}

	// check if the email provided by 'guest' is already in use by 'user' or not
	exists, err := db.CheckEmailIsAlreadyInUseByUser(userProfile.Email, ctx)
	if err != nil {
		loggerconfig.Error("Signup (service) failed to check email in use by user with error: ", err)
		return apihelpers.SendInternalServerError("Failed to check email in use by user with error: " + err.Error())
	}

	if exists {
		loggerconfig.Error("Signup (service) Email provided by guest is already in use by user")
		return apihelpers.SendErrorResponse(" Email provided by guest is already in use by user", http.StatusForbidden)
	}

	// check the count of the email so that no two users exist with same username
	count, err := db.GetGuestUsernameCount(userProfile.Email, ctx)
	userProfile.UserName = userProfile.UserName + "-" + strconv.Itoa(count+1)

	err = db.CreateUserProfile(userProfile)
	if err != nil {
		loggerconfig.Error("Signup failed to create user profile in db with error: ", err)
		return apihelpers.SendInternalServerError("Some internal server occurred! error: " + err.Error())
	}

	jwtClaims := models.Claims{
		Uuid:  userProfile.UserId,
		Email: userProfile.UserName,
		Role:  constants.RoleGuest,
		Lock:  false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
		},
	}
	token, err := middleware.GenerateToken(jwtClaims)
	if err != nil {
		loggerconfig.Error("Signup failed to generate jwt token with error: ", err)
		return apihelpers.SendInternalServerError("Unable to sign-up due to error : " + err.Error())
	}

	apiRes.Status = true
	apiRes.Data = models.AuthRes{
		AuthToken: token,
		Email:     userProfile.Email,
	}
	apiRes.Message = "auth-successful"
	return http.StatusOK, apiRes
}

func SendEmailOTP(email string) (int, apihelpers.APIRes) {
	var apiRes apihelpers.APIRes
	ctx := context.Background()

	otp, err := helpers.GenerateOTP(constants.OtpLengthForVerifyEmail)
	if err != nil {
		otp = "789451"
	}

	err = dbops.RedisSetValue(ctx, email, constants.OtpField, []byte(otp), 2*time.Minute)
	if err != nil {
		return apihelpers.SendInternalServerError("Redis error in setting otp-val with error: " + err.Error())
	}

	content := "Your otp is : " + otp
	loggerconfig.Info("OTP for email: ", email, " is ", otp)

	err = helpers.SendEmail(email, content)
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

	err = db.MarkUserEmailAsVerified(email)
	if err != nil {
		loggerconfig.Error("Failed to mark email as verified in db with error: " + err.Error())
		return apihelpers.SendInternalServerError("Failed to mark email as verified in db with error: " + err.Error())
	}

	err = helpers.SendNotification("", email, "SUCCESS")
	if err != nil {
		loggerconfig.Error("Failed to send notification with error: " + err.Error())
	}

	apiRes.Status = true
	apiRes.Message = constants.Success
	return http.StatusOK, apiRes
}

func Login(payload models.Login) (int, apihelpers.APIRes) {
	var apiRes apihelpers.APIRes

	var jwtClaims models.Claims
	switch payload.Role {
	case constants.RoleUser:
		userProfile, err := db.GetUserProfileByEmail(payload.EmailId)
		if err != nil {
			loggerconfig.Error("Login failed to fetch user profile from db with error: ", err)
			return apihelpers.SendInternalServerError("Failed to fetch user profile from db with error: " + err.Error())
		}
		if userProfile.Password != payload.Password {
			apiRes.Status = false
			apiRes.Message = "Invalid credentials"
			return http.StatusUnauthorized, apiRes
		}
		jwtClaims = models.Claims{
			Uuid:  userProfile.UserId,
			Email: userProfile.UserName,
			Role:  constants.RoleGuest,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
			},
		}
	case constants.RoleGuest:
		userProfile, err := db.GetUserProfileByUsername(payload.Username)
		if err != nil {
			loggerconfig.Error("Login failed to fetch guest profile from db with error: ", err)
			return apihelpers.SendInternalServerError("Failed to fetch guest profile from db with error: " + err.Error())
		}
		if userProfile.Password != payload.Password {
			apiRes.Status = false
			apiRes.Message = "Invalid credentials"
			return http.StatusUnauthorized, apiRes
		}
		jwtClaims = models.Claims{
			Uuid:  userProfile.UserId,
			Email: userProfile.UserName,
			Role:  constants.RoleGuest,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1)).Unix(),
			},
		}
	}

	token, err := middleware.GenerateToken(jwtClaims)
	if err != nil {
		loggerconfig.Error("Signup failed to generate jwt token with error: ", err)
		return apihelpers.SendInternalServerError("Unable to sign-up due to error : " + err.Error())
	}

	apiRes.Status = true
	apiRes.Message = constants.Success
	apiRes.Data = models.AuthRes{
		AuthToken: token,
		Email:     jwtClaims.Email,
	}
	return http.StatusOK, apiRes
}
