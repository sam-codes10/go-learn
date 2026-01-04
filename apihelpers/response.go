package apihelpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIRes struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendInternalServerError(message string) (int, APIRes) {
	var apiRes APIRes
	apiRes.Status = false
	apiRes.Message = message
	return http.StatusInternalServerError, apiRes
}

func SendBadRequest(c *gin.Context, message string) {
	var apiRes APIRes
	apiRes.Status = false
	apiRes.Message = message
	CustomResponse(c, http.StatusBadRequest, apiRes)
}

func CustomResponse(c *gin.Context, code int, data interface{}, optionalParams ...interface{}) {

	//var optionalParamsStr string
	// if len(optionalParams) > 0 {
	// 	optionalParamsStr = fmt.Sprintf(" optionalParams: %v", optionalParams)
	// }

	//logrus.Info("CustomResponse ", optionalParamsStr, " code: ", code, " | requestId:", requestId, " | clientId:", clientId, " | platform:", platform, " | clientVersion:", clientVersion)

	// Send the JSON response
	c.JSON(code, data)
}

func SendErrorResponse(message string, code int) (int, APIRes) {
	var apiRes APIRes
	apiRes.Status = false
	apiRes.Message = message
	return code, apiRes
}
