package usecase

import (
	"pustaka-api/model/mysqlmodels"
	"pustaka-api/request"
	"pustaka-api/repository"
)

type OrderDetailUsecase interface {
	FindAll(req request.OrderDetailReq) (interface{}, error)
	FindBy(req request.OrderDetailReq) (interface{}, error)
	Create(req request.OrderDetailReq) (interface{}, error)
	Update(req request.OrderDetailReq) (interface{}, error)
	Delete(req request.OrderDetailReq) (interface{}, error)
}

type orderDetailUsecase struct {
	orderDetailRepository repository.OrderDetailRepository
}

func NewOrderDetailUsecase(orderDetailRepository repository.OrderDetailRepository) OrderDetailUsecase {
	return &orderDetailUsecase{orderDetailRepository: orderDetailRepository}
}

func (srv *orderDetailUsecase) FindAll(req request.OrderDetailReq) (interface{}, error) {
	var orderDetails []*mysqlmodels.OrderDetail
	criteria := make(map[string]interface{})
	orderDetails, err := srv.orderDetailRepository.FindAll(criteria, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return orderDetails, nil
}

func (srv *orderDetailUsecase) FindBy(req request.OrderDetailReq) (interface{}, error) {
	var orderDetail *mysqlmodels.OrderDetail
	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	orderDetail, err := srv.orderDetailRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	return orderDetail, nil
}

func (srv *orderDetailUsecase) Create(req request.OrderDetailReq) (interface{}, error) {

	model := new(mysqlmodels.OrderDetail)
	model.ProductID = req.ProductID
	model.OrderID = req.OrderID
	model.OrderPrice = req.OrderPrice
	model.Qty = req.Qty

	OrderDetail, err := srv.orderDetailRepository.Create(model)

	if err != nil {
		return nil, err
	}
	return OrderDetail, nil
}

func (srv *orderDetailUsecase) Update(req request.OrderDetailReq) (interface{}, error) {

	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	model, err := srv.orderDetailRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	model.ProductID = req.ProductID
	model.OrderPrice = req.OrderPrice
	model.Qty = req.Qty

	orderDetail, err := srv.orderDetailRepository.Update(model)
	if err != nil {
		return nil, err
	}
	return orderDetail, nil
}

func (srv *orderDetailUsecase) Delete(req request.OrderDetailReq) (interface{}, error) {

	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	model, err := srv.orderDetailRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	err = srv.orderDetailRepository.Delete(model)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "success delete",
	}, nil
}
