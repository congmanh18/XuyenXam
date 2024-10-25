package mapper

import (
	"time"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"github.com/congmanh18/XuyenXam/ArtistService/model/req"
	"github.com/congmanh18/lucas-core/pointer"
	"github.com/congmanh18/lucas-core/record"
)

func TransformUpdateReqToEntity(req req.ArtistUpdateReq) *entity.Artist {
	return &entity.Artist{
		BaseEntity: record.BaseEntity{
			UpdatedAt: pointer.Time(time.Now()),
		},
		Firstname:    pointer.String(req.Firstname),
		Lastname:     pointer.String(req.Lastname),
		Username:     pointer.String(req.Username),
		Email:        pointer.String(req.Email),
		Biography:    pointer.String(req.Biography),
		Styles:       pointer.Slice(req.Styles),
		Availability: &req.Availability,
		Location:     &req.Location,
		SocialLinks:  pointer.Slice(req.SocialLinks),
	}
}
