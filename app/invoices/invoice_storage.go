package invoice

import "gorm.io/gorm"

type InvoiceStorage interface {
	GetCount(condition interface{}) (int64, error)

	/** Invoice Items */
	CreateInvoiceItem(orderItem InvoiceItemModel) (*InvoiceItemModel, error)
	GetInvoiceItems(orderId uint) ([]InvoiceItemModel, error)
}

type InvoiceStorageImpl struct {
	db *gorm.DB
}

func NewInvoiceStorage(db *gorm.DB) *InvoiceStorageImpl {
	return &InvoiceStorageImpl{db: db}
}

// func (storage *InvoiceStorageImpl) GetCount(condition interface{}) (int64, error) {
// 	var totalRecords int64
// 	query := storage.db.Model(&InvoiceModel{})
// 	if condition != nil {
// 		query = query.Where(condition)
// 	}
// 	result := query.Count(&totalRecords)
// 	if result.Error != nil {
// 		return 0, result.Error
// 	}
// 	return totalRecords, nil
// }

// func (storage *InvoiceStorageImpl) GetAll(page, pageSize int) ([]InvoiceModel, error) {
// 	var Invoices []InvoiceModel
// 	offset := (page - 1) * pageSize
// 	result := storage.db.Offset(offset).Limit(pageSize).Find(&Invoices)
// 	fmt.Println(result.Error)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return Invoices, nil
// }

// func (storage *InvoiceStorageImpl) GetById(id uint) (*InvoiceModel, error) {
// 	var Invoice *InvoiceModel
// 	result := storage.db.First(&Invoice, id)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, ErrInvoiceNotFound
// 		}
// 		return nil, result.Error
// 	}
// 	return Invoice, nil
// }

// func (storage *InvoiceStorageImpl) GetByInvoicename(Invoicename string) (*InvoiceModel, error) {
// 	var Invoice *InvoiceModel
// 	result := storage.db.Where("name = ?", Invoicename).First(&Invoice)
// 	if result.Error != nil {
// 		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 			return nil, ErrInvoiceNotFound
// 		}
// 		return nil, result.Error
// 	}
// 	return Invoice, nil
// }

// func (storage *InvoiceStorageImpl) Create(invoice InvoiceModel) (*InvoiceModel, error) {
// 	result := storage.db.Create(&invoice)
// 	if err := result.Error; err != nil {
// 		return nil, err
// 	}
// 	return &invoice, nil
// }

// func (storage *InvoiceStorageImpl) Update(Invoice InvoiceModel) error {
// 	result := storage.db.Save(Invoice)
// 	return result.Error
// }

// func (storage *InvoiceStorageImpl) Delete(id uint) error {
// 	result := storage.db.Delete(&InvoiceModel{}, id)
// 	return result.Error
// }

// /** Invoice items */
// func (storage *InvoiceStorageImpl) CreateInvoiceItem(orderItem InvoiceItemModel) (*InvoiceItemModel, error) {
// 	result := storage.db.Create(&orderItem)
// 	if err := result.Error; err != nil {
// 		return nil, err
// 	}
// 	return &orderItem, nil
// }

// func (storage *InvoiceStorageImpl) GetInvoiceItems(orderId uint) ([]InvoiceItemModel, error) {
// 	var orderItems []InvoiceItemModel
// 	result := storage.db.Find(&orderItems)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return orderItems, nil
// }
