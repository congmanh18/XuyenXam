package provider

import (
	"github.com/congmanh18/XuyenXam/OrderService/model"
	"github.com/congmanh18/lucas-core/db/postgresql"
)

type AppProvider struct {
	Postgres *postgresql.Database
}

func NewAppProvider(config model.ServiceConfig) *AppProvider {
	return &AppProvider{
		Postgres: ProvidePostgres(config),
	}
}
