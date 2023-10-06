package handler

import (
	"net/http"

	_"pustaka-api/helper"
	"pustaka-api/request"
	"pustaka-api/responses"
	"pustaka-api/usecase"

	"github.com/gin-gonic/gin"
)

type orderHandler struct {
	uc usecase.OrderUsecase
}

func NewOrderHandler(uc usecase.OrderUsecase) *orderHandler {
	return &orderHandler{uc: uc}
}

func (h *orderHandler) GetOrdersHandler(ctx *gin.Context) {
	var req request.OrderReq

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
	// req.CustomerID = int(authCust)
	res, err := h.uc.FindAll(req)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *orderHandler) GetOrderHandler(ctx *gin.Context) {

	var req request.OrderReq
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

func (h *orderHandler) PostOrderHandler(ctx *gin.Context) {
	var order request.OrderReq

	err := ctx.ShouldBindJSON(&order)
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

	res, err := h.uc.Create(order)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *orderHandler) UpdateOrderHandler(ctx *gin.Context) {
	var order request.OrderReq

	err := ctx.ShouldBindUri(&order)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	err = ctx.ShouldBindJSON(&order)
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

	res, err := h.uc.Update(order)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}

func (h *orderHandler) DeleteOrderHandler(ctx *gin.Context) {
	var order request.OrderReq

	err := ctx.ShouldBindUri(&order)
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

	res, err := h.uc.Delete(order)
	if err != nil {
		responses.NewErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	responses.NewSuccessResponse(ctx, res)
}
