package mapper

import (
	"time"

	"github.com/congmanh18/XuyenXam/ArtistService/entity"
	"github.com/congmanh18/XuyenXam/ArtistService/model/req"
	"github.com/congmanh18/lucas-core/pointer"
	"github.com/congmanh18/lucas-core/record"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TransformCreateReqToEntity(req req.ArtistCreateReq) *entity.Artist {
	return &entity.Artist{
		BaseEntity: record.BaseEntity{
			ID:        pointer.String(primitive.NewObjectID().String()),
			CreatedAt: pointer.Time(time.Now()),
			UpdatedAt: pointer.Time(time.Now()),
		},
		Firstname:   pointer.String(req.Firstname),
		Lastname:    pointer.String(req.Lastname),
		Email:       pointer.String(req.Email),
		PhoneNumber: pointer.String(req.PhoneNumber),
		Biography:   pointer.String(req.Biography),
		SocialLinks: pointer.Slice(req.SocialLinks),
	}
}
