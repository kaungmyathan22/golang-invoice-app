package order

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type OrderStorage interface {
	GetCount(condition interface{}) (int64, error)
	GetAll(page, pageSize int) ([]OrderModel, error)
	GetById(id uint) (*OrderModel, error)
	Create(Order OrderModel) (*OrderModel, error)
	Update(Order OrderModel) error
	Delete(id uint) error

	/** Order Items */
	CreateOrderItem(orderItem OrderItemModel) (*OrderItemModel, error)
	GetOrderItems(orderId uint) ([]OrderItemModel, error)
}

type OrderStorageImpl struct {
	db *gorm.DB
}

func NewOrderStorage(db *gorm.DB) *OrderStorageImpl {
	return &OrderStorageImpl{db: db}
}

func (storage *OrderStorageImpl) GetCount(condition interface{}) (int64, error) {
	var totalRecords int64
	query := storage.db.Model(&OrderModel{})
	if condition != nil {
		query = query.Where(condition)
	}
	result := query.Count(&totalRecords)
	if result.Error != nil {
		return 0, result.Error
	}
	return totalRecords, nil
}

func (storage *OrderStorageImpl) GetAll(page, pageSize int) ([]OrderModel, error) {
	var Orders []OrderModel
	offset := (page - 1) * pageSize
	result := storage.db.Offset(offset).Limit(pageSize).Find(&Orders)
	fmt.Println(result.Error)
	if result.Error != nil {
		return nil, result.Error
	}
	return Orders, nil
}

func (storage *OrderStorageImpl) GetById(id uint) (*OrderModel, error) {
	var Order *OrderModel
	result := storage.db.First(&Order, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, result.Error
	}
	return Order, nil
}

func (storage *OrderStorageImpl) GetByOrdername(Ordername string) (*OrderModel, error) {
	var Order *OrderModel
	result := storage.db.Where("name = ?", Ordername).First(&Order)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrOrderNotFound
		}
		return nil, result.Error
	}
	return Order, nil
}

func (storage *OrderStorageImpl) Create(order OrderModel) (*OrderModel, error) {
	result := storage.db.Create(&order)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (storage *OrderStorageImpl) Update(Order OrderModel) error {
	result := storage.db.Save(Order)
	return result.Error
}

func (storage *OrderStorageImpl) Delete(id uint) error {
	result := storage.db.Delete(&OrderModel{}, id)
	return result.Error
}

/** Order items */
func (storage *OrderStorageImpl) CreateOrderItem(orderItem OrderItemModel) (*OrderItemModel, error) {
	result := storage.db.Create(&orderItem)
	if err := result.Error; err != nil {
		return nil, err
	}
	return &orderItem, nil
}
func (storage *OrderStorageImpl) GetOrderItems(orderId uint) ([]OrderItemModel, error) {
	var orderItems []OrderItemModel
	result := storage.db.Find(&orderItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderItems, nil
}
