package handler

import (
	"log"
	"net/http"

	"pustaka-api/request"
	"pustaka-api/responses"
	"pustaka-api/usecase"

	"github.com/gin-gonic/gin"
)

type customerHandler struct {
	uc usecase.CustomerUsecase
}

func NewCustomerHandler(uc usecase.CustomerUsecase) *customerHandler {
	return &customerHandler{uc: uc}
}

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"name": "Alvin",
	})
}

func (h *customerHandler) GetcustomersHandler(ctx *gin.Context) {
	var req request.CustomerReq
	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	res, err := h.uc.FindAll(req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *customerHandler) GetcustomerHandler(ctx *gin.Context) {

	var req request.CustomerReq
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	log.Println(req)
	res, err := h.uc.FindBy(req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *customerHandler) PostcustomerHandler(ctx *gin.Context) {
	var customer request.CustomerReq

	err := ctx.ShouldBindJSON(&customer)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := h.uc.Create(customer)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *customerHandler) UpdatecustomerHandler(ctx *gin.Context) {
	var customer request.CustomerReq

	err := ctx.ShouldBindUri(&customer)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&customer)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := h.uc.Update(customer)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *customerHandler) DeletecustomerHandler(ctx *gin.Context) {
	var customer request.CustomerReq

	err := ctx.ShouldBindUri(&customer)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := h.uc.Delete(customer)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}
