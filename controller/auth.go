package controller

import (
	"auth-service/apihelpers"
	"auth-service/models"
	"auth-service/service"

	"github.com/gin-gonic/gin"
)

// @Tags Auth
// @Summary User Signup
// @Description API for user signup
// @Param body models.SignUp true "Signup Payload"
// @Success 200 {object} apihelpers.APIRes
// @Router /v1/auth/signup [post]
func SignUp(c *gin.Context) {
	var payload models.SignUp
	
	if err := c.ShouldBindJSON(&payload); err != nil {
		apihelpers.SendBadRequest(c, "invalid payload")
		return
	}

	code, resp := service.Signup(payload)
	apihelpers.CustomResponse(c, code, resp)
}
