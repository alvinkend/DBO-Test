package repository

import (
	"pustaka-api/model/mysqlmodels"
	"pustaka-api/helper"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll(criteria map[string]interface{}, page, size int) ([]*mysqlmodels.Order, error)
	FindBy(criteria map[string]interface{}) (*mysqlmodels.Order, error)
	Create(order *mysqlmodels.Order) (*mysqlmodels.Order, error)
	Update(order *mysqlmodels.Order) (*mysqlmodels.Order, error)
	Delete(order *mysqlmodels.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (srv *orderRepository) FindAll(criteria map[string]interface{}, page, size int) ([]*mysqlmodels.Order, error) {
	var order []*mysqlmodels.Order

	if page == 0 || size == 0 {
		page, size = -1, -1
	}

	limit, offset := helper.GetLimitOffset(page, size)

	if res := srv.db.Where(criteria).Offset(offset).Limit(limit).Find(&order); res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}

func (srv *orderRepository) FindBy(criteria map[string]interface{}) (*mysqlmodels.Order, error) {
	var order *mysqlmodels.Order

	/*
		Documentation from gorm why use .First when query findby:
		GORM V2 only returns ErrRecordNotFound when you are querying with methods First, Last, Take
		which is expected to return some result, and we have also removed method RecordNotFound in V2
	*/
	if res := srv.db.Where(criteria).First(&order); res.Error != nil {
		return nil, res.Error
	}
	return order, nil
}

func (srv *orderRepository) Create(model *mysqlmodels.Order) (*mysqlmodels.Order, error) {

	db := srv.db.Create(&model)
	if err := db.Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (srv *orderRepository) Update(model *mysqlmodels.Order) (*mysqlmodels.Order, error) {
	err := srv.db.Save(&model).Error
	return model, err
}

func (srv *orderRepository) Delete(model *mysqlmodels.Order) error {
	err := srv.db.Delete(&model).Error
	return err
}
