package mapper

import (
	"github.com/congmanh18/XuyenXam/OrderService/entity"
	"github.com/congmanh18/XuyenXam/OrderService/model/req"
	"github.com/congmanh18/lucas-core/pointer"
	"github.com/congmanh18/lucas-core/record"
	"github.com/google/uuid"
)

func TransformOrderCreateReqToEntity(req req.OrderCreateReq) *entity.Order {
	return &entity.Order{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		CustomerID: req.CustomerID,
		ArtistID:   req.ArtistID,
		Status:     req.Status,
		TotalPrice: req.TotalPrice,
		Items:      TransformOrderItemsCreateReqToEntity(req.Items),
	}
}

func TransformOrderItemsCreateReqToEntity(items []req.OrderItemCreateReq) []entity.OrderItem {
	var orderItems []entity.OrderItem
	for _, item := range items {
		orderItems = append(orderItems, entity.OrderItem{
			BaseEntity: record.BaseEntity{
				ID: pointer.String(uuid.New().String()),
			},
			ProductName: item.ProductName,
			Quantity:    item.Quantity,
			Price:       item.Price,
		})
	}
	return orderItems
}
