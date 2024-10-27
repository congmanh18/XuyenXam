package mapper

import (
	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"github.com/congmanh18/XuyenXam/ArtistService/model/req"
	"github.com/congmanh18/lucas-core/pointer"
	"github.com/congmanh18/lucas-core/record"
)

// Thừa ???
func TransformBaseReqToEntity(req req.BaseReq) *entity.Artist {
	return &entity.Artist{
		BaseEntity: record.BaseEntity{
			ID: pointer.String(req.ID),
		},
	}
}
