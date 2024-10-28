package provider

import (
	"fmt"

	"github.com/congmanh18/XuyenXam/ProductService/model"
	"github.com/congmanh18/lucas-core/db/postgresql"
)

func ProvidePostgres(config model.ServiceConfig) *postgresql.Database {
	fmt.Println("Connecting to PostgreSQL")
	return postgresql.New(postgresql.Connection{
		Host:     config.DBHost,
		Port:     config.DBPort,
		Database: config.DBName,
		User:     config.DBUser,
		Password: config.DBPwd,
		SSLMode:  postgresql.Disable,
	})
}
