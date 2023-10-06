package usecase

import (
	"pustaka-api/model/mysqlmodels"
	"pustaka-api/repository"
	"pustaka-api/request"
)

type OrderUsecase interface {
	FindAll(req request.OrderReq) (interface{}, error)
	FindBy(req request.OrderReq) (interface{}, error)
	Create(req request.OrderReq) (interface{}, error)
	Update(req request.OrderReq) (interface{}, error)
	Delete(req request.OrderReq) (interface{}, error)
}

type orderUsecase struct {
	orderRepository repository.OrderRepository
}

func NewOrderUsecase(orderRepository repository.OrderRepository) OrderUsecase {
	return &orderUsecase{orderRepository: orderRepository}
}

func (srv *orderUsecase) FindAll(req request.OrderReq) (interface{}, error) {
	var orders []*mysqlmodels.Order
	criteria := make(map[string]interface{})
	//criteria["customer_id"] = req.CustomerID
	orders, err := srv.orderRepository.FindAll(criteria, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (srv *orderUsecase) FindBy(req request.OrderReq) (interface{}, error) {
	var order *mysqlmodels.Order
	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	order, err := srv.orderRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (srv *orderUsecase) Create(req request.OrderReq) (interface{}, error) {

	model := new(mysqlmodels.Order)
	model.CustomerID = req.CustomerID
	model.OrderDate = req.OrderDate
	model.TotalAmount = req.TotalAmount

	order, err := srv.orderRepository.Create(model)

	if err != nil {
		return nil, err
	}
	return order, nil
}

func (srv *orderUsecase) Update(req request.OrderReq) (interface{}, error) {

	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	model, err := srv.orderRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	model.OrderDate = req.OrderDate
	model.TotalAmount = req.TotalAmount

	customer, err := srv.orderRepository.Update(model)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (srv *orderUsecase) Delete(req request.OrderReq) (interface{}, error) {

	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	model, err := srv.orderRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	err = srv.orderRepository.Delete(model)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "success delete",
	}, nil
}
