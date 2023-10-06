package handler

import (
	"net/http"

	"pustaka-api/request"
	"pustaka-api/responses"
	"pustaka-api/usecase"
	_"pustaka-api/helper"

	"github.com/gin-gonic/gin"
)


type orderDetailHandler struct {
	uc usecase.OrderDetailUsecase
}

func NewOrderDetailHandler(uc usecase.OrderDetailUsecase) *orderDetailHandler {
	return &orderDetailHandler{uc: uc}
}

func (h *orderDetailHandler) GetOrderDetailsHandler(ctx *gin.Context) {
	var req request.OrderDetailReq

	err := ctx.ShouldBindQuery(&req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// authCust, err := helper.GetAuthenticatedUser(ctx.Request)
	// if err != nil {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }

	// if authCust == 0 {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }

	res, err := h.uc.FindAll(req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *orderDetailHandler) GetOrderDetailHandler(ctx *gin.Context) {

	var req request.OrderDetailReq
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// authCust, err := helper.GetAuthenticatedUser(ctx.Request)
	// if err != nil {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }

	// if authCust == 0 {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }
	
	res, err := h.uc.FindBy(req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *orderDetailHandler) PostOrderDetailHandler(ctx *gin.Context) {
	var orderDetail request.OrderDetailReq

	err := ctx.ShouldBindJSON(&orderDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// authCust, err := helper.GetAuthenticatedUser(ctx.Request)
	// if err != nil {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }

	// if authCust == 0 {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }

	res, err := h.uc.Create(orderDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *orderDetailHandler) UpdateOrderDetailHandler(ctx *gin.Context) {
	var orderDetail request.OrderDetailReq

	err := ctx.ShouldBindUri(&orderDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&orderDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// authCust, err := helper.GetAuthenticatedUser(ctx.Request)
	// if err != nil {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }

	// if authCust == 0 {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }

	res, err := h.uc.Update(orderDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *orderDetailHandler) DeleteOrderDetailHandler(ctx *gin.Context) {
	var orderDetail request.OrderDetailReq

	err := ctx.ShouldBindUri(&orderDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	// authCust, err := helper.GetAuthenticatedUser(ctx.Request)
	// if err != nil {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }

	// if authCust == 0 {
	// 	responses.NewErrorResponse(ctx, http.StatusForbidden, err)
	// 	return
	// }

	res, err := h.uc.Delete(orderDetail)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}