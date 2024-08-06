package order

func (model *OrderModel) ToEntity() *OrderEntity {
	return &OrderEntity{
		ID:              model.ID,
		CustomerName:    model.CustomerName,
		OrderStatus:     model.OrderStatus.String(),
		CustomerPhoneNo: model.CustomerPhoneNo,
		BillingAddress:  model.BillingAddress,
		OrderNo:         model.OrderNo,
		ShippingAddress: model.ShippingAddress,
		ShippingCosts:   model.ShippingCosts,
		SubTotal:        model.SubTotal,
		Total:           model.Total,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		DeletedAt:       model.DeletedAt,
	}
}

func (model *OrderItemModel) ToEntity() *OrderItemEntity {
	return &OrderItemEntity{
		ID:        model.ID,
		ProductId: model.ProductId,
		OrderId:   model.OrderId,
		Quantity:  model.Quantity,
		Total:     model.Total,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		DeletedAt: model.DeletedAt,
	}
}

func (entity *OrderItemEntity) ToModel() *OrderItemModel {
	return &OrderItemModel{
		ID:        entity.ID,
		ProductId: entity.ProductId,
		OrderId:   entity.OrderId,
		Total:     entity.Total,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
}

func (entity *OrderEntity) ToModel() *OrderModel {
	return &OrderModel{
		ID:              entity.ID,
		OrderStatus:     OrderStatus(entity.OrderStatus),
		CustomerName:    entity.CustomerName,
		CustomerPhoneNo: entity.CustomerPhoneNo,
		BillingAddress:  entity.BillingAddress,
		ShippingAddress: entity.ShippingAddress,
		ShippingCosts:   entity.ShippingCosts,
		SubTotal:        entity.SubTotal,
		Total:           entity.Total,
		CreatedAt:       entity.CreatedAt,
		UpdatedAt:       entity.UpdatedAt,
		DeletedAt:       entity.DeletedAt,
	}
}
