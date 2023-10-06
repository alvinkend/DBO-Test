package usecase

import (
	"pustaka-api/model/mysqlmodels"
	"pustaka-api/repository"
	"pustaka-api/request"

	"golang.org/x/crypto/bcrypt"
)

type CustomerUsecase interface {
	FindAll(req request.CustomerReq) (interface{}, error)
	FindBy(req request.CustomerReq) (interface{}, error)
	Create(req request.CustomerReq) (interface{}, error)
	Update(req request.CustomerReq) (interface{}, error)
	Delete(req request.CustomerReq) (interface{}, error)
}

type customerUsecase struct {
	customerRepository repository.CustomerRepository
}

func NewCustomerUsecase(customerRepository repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{customerRepository: customerRepository}
}

func (srv *customerUsecase) FindAll(req request.CustomerReq) (interface{}, error) {
	var customers []*mysqlmodels.Customer
	criteria := make(map[string]interface{})
	customers, err := srv.customerRepository.FindAll(criteria, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (srv *customerUsecase) FindBy(req request.CustomerReq) (interface{}, error) {
	var customer *mysqlmodels.Customer
	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	customer, err := srv.customerRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (srv *customerUsecase) Create(req request.CustomerReq) (interface{}, error) {

	model := new(mysqlmodels.Customer)
	model.CustomerName = req.CustomerName
	model.Email = req.Email
	model.ShippingAddres = req.ShippingAddres

	bytePassword := []byte(req.Password)
	cusPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	model.Password = string(cusPassword)

	customer, err := srv.customerRepository.Create(model)

	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (srv *customerUsecase) Update(req request.CustomerReq) (interface{}, error) {

	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	model, err := srv.customerRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	model.CustomerName = req.CustomerName
	model.Email = req.Email
	model.ShippingAddres = req.ShippingAddres

	customer, err := srv.customerRepository.Update(model)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (srv *customerUsecase) Delete(req request.CustomerReq) (interface{}, error) {

	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	model, err := srv.customerRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	err = srv.customerRepository.Delete(model)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "success delete",
	}, nil
}
