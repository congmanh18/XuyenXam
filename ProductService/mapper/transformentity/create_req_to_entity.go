package mapper

import (
	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/model/req"
	"github.com/congmanh18/lucas-core/pointer"
	"github.com/congmanh18/lucas-core/record"
	"github.com/google/uuid"
)

func TransformProductCreateReqToEntity(req req.ProductCreateReq) *entity.Product {
	return &entity.Product{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(uuid.New().String()),
		},
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
}

func TransformProductCategoryCreateReqToEntity(req req.ProductCategoryCreateReq) *entity.ProductCategory {
	return &entity.ProductCategory{
		ProductID:  req.ProductID,
		CategoryID: req.CategoryID,
	}
}

func TransformCategoryCreateReqToEntity(req req.CategoryCreateReq) *entity.Category {
	return &entity.Category{
		Name:        req.Name,
		Description: req.Description,
	}
}
