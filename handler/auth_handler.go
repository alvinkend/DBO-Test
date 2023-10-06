package handler

import (
	"net/http"

	"pustaka-api/request"
	"pustaka-api/responses"
	"pustaka-api/usecase"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	uc usecase.AuthUsecase
}

func NewAuthHandler(uc usecase.AuthUsecase) *authHandler {
	return &authHandler{uc: uc}
}

func (h *authHandler) LoginHandler(ctx *gin.Context) {
	var req request.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := h.uc.LoginUsecase(req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}