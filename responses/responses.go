package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
	Status    int    `json:"status"`
}

type TokenResponse struct {
	AccessToken string  `json:"access_token"`
	TokenType   string  `json:"token_type"`
	Role   		string  `json:"role"`
	ExpiredIn   float64 `json:"expired_in"`
	ExpiredAt   int64   `json:"expired_at"`
}

func NewErrorResponse(c *gin.Context, code int, err error) {

	response := new(ResponseError)
	response.Status = code
	response.Message = err.Error()
	response.ErrorCode = code
	c.JSON(code, &response)

	return
}

func NewSuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
	return
}
