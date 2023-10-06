package handler

import (
	"net/http"

	"pustaka-api/request"
	"pustaka-api/responses"
	"pustaka-api/usecase"

	"github.com/gin-gonic/gin"
)


type customerDetailHandler struct {
	uc usecase.CustomerDetailUsecase
}

func NewCustomerDetailHandler(uc usecase.CustomerDetailUsecase) *customerDetailHandler {
	return &customerDetailHandler{uc: uc}
}

func (h *customerDetailHandler) GetCustomerDetailsHandler(ctx *gin.Context) {
	var req request.CustomerDetailReq
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

func (h *customerDetailHandler) GetCustomerDetailHandler(ctx *gin.Context) {

	var req request.CustomerDetailReq
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	res, err := h.uc.FindBy(req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *customerDetailHandler) PostCustomerDetailHandler(ctx *gin.Context) {
	var customerDetail request.CustomerDetailReq

	err := ctx.ShouldBindJSON(&customerDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := h.uc.Create(customerDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *customerDetailHandler) UpdateCustomerDetailHandler(ctx *gin.Context) {
	var customerDetail request.CustomerDetailReq

	err := ctx.ShouldBindUri(&customerDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&customerDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := h.uc.Update(customerDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *customerDetailHandler) DeleteCustomerDetailHandler(ctx *gin.Context) {
	var customerDetail request.CustomerDetailReq

	err := ctx.ShouldBindUri(&customerDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	res, err := h.uc.Delete(customerDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}