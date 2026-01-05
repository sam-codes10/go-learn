package controller

import (
	"auth-service/apihelpers"
	"auth-service/models"
	"auth-service/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @Tags Auth
// @Summary User Signup
// @Description API for user signup
// @Param body body models.SignUp true "Signup Payload"
// @Success 200 {object} apihelpers.APIRes{data=models.SignUpRes}
// @Router /v1/auth/signup [post]
func SignUp(c *gin.Context) {
	var payload models.SignUp

	if err := c.ShouldBindJSON(&payload); err != nil {
		apihelpers.SendBadRequest(c, "invalid payload")
		return
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		apihelpers.SendBadRequest(c, "validation error: "+err.Error())
		return
	}

	code, resp := service.Signup(payload)
	apihelpers.CustomResponse(c, code, resp)
}

// @Tags Auth
// @Summary send email otp
// @Description API for sending otp by email
// @Param email query string true "email" default(satish.sund3r@gmail.com)
// @Success 200 {object} apihelpers.APIRes
// @Router /v1/auth/send-email-otp [get]
func SendEmailOTP(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		apihelpers.SendBadRequest(c, "Email field is empty")
	}

	code, resp := service.SendEmailOTP(email)
	apihelpers.CustomResponse(c, code, resp)
}

// @Tags Auth
// @Summary verify email otp
// @Description API for verifying otp by email
// @Param email query string true "email" default(satish.sund3r@gmail.com)
// @Param otp query string true "otp" defualt(123456)
// @Success 200 {object} apihelpers.APIRes
// @Router /v1/auth/verify-email-otp [get]
func VerifyEmailOtp(c *gin.Context) {
	email := c.Query("email")
	otp := c.Query("otp")
	if email == "" || otp == "" {
		apihelpers.SendBadRequest(c, "Email or otp field is empty")
	}

	code, resp := service.VerifyEmailOtp(email, otp)
	apihelpers.CustomResponse(c, code, resp)
}

// @Tags Auth
// @Summary User Login
// @Description API for user login
// @Param body body models.Login true "Login Payload"
// @Success 200 {object} apihelpers.APIRes
// @Router /v1/auth/login [post]
func Login(c *gin.Context) {
	var payload models.Login
	if err := c.ShouldBindJSON(&payload); err != nil {
		apihelpers.SendBadRequest(c, "invalid payload")
		return
	}
	code, resp := service.Login(payload)
	apihelpers.CustomResponse(c, code, resp)
}
