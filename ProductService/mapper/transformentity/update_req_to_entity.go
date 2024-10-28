package mapper

import (
	"github.com/congmanh18/XuyenXam/ProductService/entity"
	"github.com/congmanh18/XuyenXam/ProductService/model/req"
)

func TransformProductUpdateReqToEntity(req req.ProductUpdateReq) *entity.Product {
	return &entity.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}
}

func TransformProductCategoryUpdateReqToEntity(req req.ProductCategoryUpdateReq) *entity.ProductCategory {
	return &entity.ProductCategory{
		ProductID:  req.ProductID,
		CategoryID: req.CategoryID,
	}
}

func TransformCategoryUpdateReqToEntity(req req.CategoryUpdateReq) *entity.Category {
	return &entity.Category{
		Name:        req.Name,
		Description: req.Description,
	}
}
