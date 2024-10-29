package mapper

import (
	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/pointer"
	"github.com/congmanh18/lucas-core/record"
	"github.com/google/uuid"
)

func TransformProductReqToEntity(req req.ProductReq) *entity.Product {
	return &entity.Product{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
}

func TransformProductCategoryReqToEntity(req req.ProductCategoryReq) *entity.ProductCategory {
	return &entity.ProductCategory{
		ProductID:  req.ProductID,
		CategoryID: req.CategoryID,
	}
}

func TransformCategoryReqToEntity(req req.CategoryReq) *entity.Category {
	return &entity.Category{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Name:        req.Name,
		Description: req.Description,
	}
}

func TransformImageReqToEntity(req req.ImageReq) *entity.Image {
	return &entity.Image{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		ProductID: req.ProductID,
		URL:       req.ImageURL,
	}
}

func TransformInventoryReqToEntity(req req.InventoryReq) *entity.Inventory {
	return &entity.Inventory{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Quantity: req.Quantity,
	}
}
