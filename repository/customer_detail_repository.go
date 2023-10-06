package repository

import (
	"pustaka-api/model/mysqlmodels"
	"pustaka-api/helper"

	"gorm.io/gorm"
)

type CustomerDetailRepository interface {
	FindAll(criteria map[string]interface{}, page, size int) ([]*mysqlmodels.CustomerDetail, error)
	FindBy(criteria map[string]interface{}) (*mysqlmodels.CustomerDetail, error)
	Create(customer *mysqlmodels.CustomerDetail) (*mysqlmodels.CustomerDetail, error)
	Update(customer *mysqlmodels.CustomerDetail) (*mysqlmodels.CustomerDetail, error)
	Delete(customer *mysqlmodels.CustomerDetail) error
}

type customerDetailRepository struct {
	db *gorm.DB
}

func NewCustomerDetailRepository(db *gorm.DB) CustomerDetailRepository {
	return &customerDetailRepository{db}
}

func (srv *customerDetailRepository) FindAll(criteria map[string]interface{}, page, size int) ([]*mysqlmodels.CustomerDetail, error) {
	var customerDetail []*mysqlmodels.CustomerDetail

	if page == 0 || size == 0 {
		page, size = -1, -1
	}

	limit, offset := helper.GetLimitOffset(page, size)

	if res := srv.db.Where(criteria).Offset(offset).Limit(limit).Find(&customerDetail); res.Error != nil {
		return nil, res.Error
	}
	return customerDetail, nil
}

func (srv *customerDetailRepository) FindBy(criteria map[string]interface{}) (*mysqlmodels.CustomerDetail, error) {
	var customerDetail *mysqlmodels.CustomerDetail

	/*
		Documentation from gorm why use .First when query findby:
		GORM V2 only returns ErrRecordNotFound when you are querying with methods First, Last, Take
		which is expected to return some result, and we have also removed method RecordNotFound in V2
	*/
	if res := srv.db.Where(criteria).First(&customerDetail); res.Error != nil {
		return nil, res.Error
	}
	return customerDetail, nil
}

func (srv *customerDetailRepository) Create(model *mysqlmodels.CustomerDetail) (*mysqlmodels.CustomerDetail, error) {

	db := srv.db.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *customerDetailRepository) Update(model *mysqlmodels.CustomerDetail) (*mysqlmodels.CustomerDetail, error) {
	err := srv.db.Save(&model).Error
	return model, err
}

func (srv *customerDetailRepository) Delete(model *mysqlmodels.CustomerDetail) error {
	err := srv.db.Delete(&model).Error
	return err
}