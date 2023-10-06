package repository

import (
	"pustaka-api/helper"
	"pustaka-api/model/mysqlmodels"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	FindAll(criteria map[string]interface{}, page, size int) ([]*mysqlmodels.Customer, error)
	FindBy(criteria map[string]interface{}) (*mysqlmodels.Customer, error)
	Create(customer *mysqlmodels.Customer) (*mysqlmodels.Customer, error)
	Update(customer *mysqlmodels.Customer) (*mysqlmodels.Customer, error)
	Delete(customer *mysqlmodels.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{db}
}

func (srv *customerRepository) FindAll(criteria map[string]interface{}, page, size int) ([]*mysqlmodels.Customer, error) {
	var customer []*mysqlmodels.Customer

	if page == 0 || size == 0 {
		page, size = -1, -1
	}

	limit, offset := helper.GetLimitOffset(page, size)
	if res := srv.db.Where(criteria).Offset(offset).Limit(limit).Find(&customer); res.Error != nil {
		return nil, res.Error
	}
	return customer, nil
}

func (srv *customerRepository) FindBy(criteria map[string]interface{}) (*mysqlmodels.Customer, error) {
	var customer *mysqlmodels.Customer

	/*
		Documentation from gorm why use .First when query findby:
		GORM V2 only returns ErrRecordNotFound when you are querying with methods First, Last, Take
		which is expected to return some result, and we have also removed method RecordNotFound in V2
	*/
	if res := srv.db.Where(criteria).First(&customer); res.Error != nil {
		return nil, res.Error
	}
	return customer, nil
}

func (srv *customerRepository) Create(model *mysqlmodels.Customer) (*mysqlmodels.Customer, error) {

	db := srv.db.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *customerRepository) Update(model *mysqlmodels.Customer) (*mysqlmodels.Customer, error) {
	err := srv.db.Save(&model).Error
	return model, err
}

func (srv *customerRepository) Delete(model *mysqlmodels.Customer) error {
	err := srv.db.Delete(&model).Error
	return err
}
