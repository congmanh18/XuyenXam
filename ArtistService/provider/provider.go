package provider

import (
	"github.com/congmanh18/XuyenXam/ArtistService/model"
	"github.com/congmanh18/lucas-core/db/mongodb"
)

type AppProvider struct {
	MongoDB *mongodb.Database
}

func NewAppProvider(config model.ServiceConfig) *AppProvider {
	return &AppProvider{
		MongoDB: ProvideMongo(config),
	}
}
