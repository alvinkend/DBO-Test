package usecase

import (
	"pustaka-api/model/mysqlmodels"
	"pustaka-api/request"
	"pustaka-api/repository"
)

type CustomerDetailUsecase interface {
	FindAll(req request.CustomerDetailReq) (interface{}, error)
	FindBy(req request.CustomerDetailReq) (interface{}, error)
	Create(req request.CustomerDetailReq) (interface{}, error)
	Update(req request.CustomerDetailReq) (interface{}, error)
	Delete(req request.CustomerDetailReq) (interface{}, error)
}

type customerDetailUsecase struct {
	customerDetailRepository repository.CustomerDetailRepository
}

func NewCustomerDetailUsecase(customerDetailRepository repository.CustomerDetailRepository) CustomerDetailUsecase {
	return &customerDetailUsecase{customerDetailRepository: customerDetailRepository}
}

func (srv *customerDetailUsecase) FindAll(req request.CustomerDetailReq) (interface{}, error) {
	var customerDetails []*mysqlmodels.CustomerDetail
	criteria := make(map[string]interface{})
	customerDetails, err := srv.customerDetailRepository.FindAll(criteria, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return customerDetails, nil
}

func (srv *customerDetailUsecase) FindBy(req request.CustomerDetailReq) (interface{}, error) {
	var customerDetail *mysqlmodels.CustomerDetail
	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	customerDetail, err := srv.customerDetailRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	return customerDetail, nil
}

func (srv *customerDetailUsecase) Create(req request.CustomerDetailReq) (interface{}, error) {

	model := new(mysqlmodels.CustomerDetail)
	model.CustomerID = req.CustomerID
	model.Gender = req.Gender
	model.DateOfBirth = req.DateOfBirth

	CustomerDetail, err := srv.customerDetailRepository.Create(model)

	if err != nil {
		return nil, err
	}
	return CustomerDetail, nil
}

func (srv *customerDetailUsecase) Update(req request.CustomerDetailReq) (interface{}, error) {

	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	model, err := srv.customerDetailRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	model.CustomerID = req.CustomerID
	model.Gender = req.Gender
	model.DateOfBirth = req.DateOfBirth

	customerDetail, err := srv.customerDetailRepository.Update(model)
	if err != nil {
		return nil, err
	}
	return customerDetail, nil
}

func (srv *customerDetailUsecase) Delete(req request.CustomerDetailReq) (interface{}, error) {

	criteria := make(map[string]interface{})
	criteria["id"] = req.ID
	model, err := srv.customerDetailRepository.FindBy(criteria)
	if err != nil {
		return nil, err
	}

	err = srv.customerDetailRepository.Delete(model)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "success delete",
	}, nil
}
