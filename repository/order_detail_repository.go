package repository

import (
	"pustaka-api/model/mysqlmodels"
	"pustaka-api/helper"

	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	FindAll(criteria map[string]interface{}, page, size int) ([]*mysqlmodels.OrderDetail, error)
	FindBy(criteria map[string]interface{}) (*mysqlmodels.OrderDetail, error)
	Create(orderDetail *mysqlmodels.OrderDetail) (*mysqlmodels.OrderDetail, error)
	Update(orderDetail *mysqlmodels.OrderDetail) (*mysqlmodels.OrderDetail, error)
	Delete(orderDetail *mysqlmodels.OrderDetail) error
}

type orderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepository{db}
}

func (srv *orderDetailRepository) FindAll(criteria map[string]interface{}, page, size int) ([]*mysqlmodels.OrderDetail, error) {
	var orderDetails []*mysqlmodels.OrderDetail

	if page == 0 || size == 0 {
		page, size = -1, -1
	}

	limit, offset := helper.GetLimitOffset(page, size)

	if res := srv.db.Where(criteria).Offset(offset).Limit(limit).Find(&orderDetails); res.Error != nil {
		return nil, res.Error
	}
	return orderDetails, nil
}

func (srv *orderDetailRepository) FindBy(criteria map[string]interface{}) (*mysqlmodels.OrderDetail, error) {
	var orderDetail *mysqlmodels.OrderDetail

	/*
		Documentation from gorm why use .First when query findby:
		GORM V2 only returns ErrRecordNotFound when you are querying with methods First, Last, Take
		which is expected to return some result, and we have also removed method RecordNotFound in V2
	*/
	if res := srv.db.Where(criteria).First(&orderDetail); res.Error != nil {
		return nil, res.Error
	}
	return orderDetail, nil
}

func (srv *orderDetailRepository) Create(model *mysqlmodels.OrderDetail) (*mysqlmodels.OrderDetail, error) {

	db := srv.db.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *orderDetailRepository) Update(model *mysqlmodels.OrderDetail) (*mysqlmodels.OrderDetail, error) {
	err := srv.db.Save(&model).Error
	return model, err
}

func (srv *orderDetailRepository) Delete(model *mysqlmodels.OrderDetail) error {
	err := srv.db.Delete(&model).Error
	return err
}
