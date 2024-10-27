package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/congmanh18/XuyenXam/ArtistService/model"
	"github.com/congmanh18/lucas-core/db/mongodb"
)

func ProvideMongo(config model.ServiceConfig) *mongodb.Database {
	fmt.Println("Connecting to MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rs, _ := mongodb.NewDatabase(&mongodb.Connection{
		Host:           config.DBHost,
		Port:           config.DBPort,
		User:           config.DBUser,
		Password:       config.DBPwd,
		Database:       config.DBName,
		AuthSource:     config.AuthSource,
		ReplicaSet:     config.ReplicaSet,
		MaxPoolSize:    config.MaxPoolSize,
		MinPoolSize:    config.MinPoolSize,
		ConnectTimeout: config.ConnectTimeout,
		SocketTimeout:  config.SocketTimeout,
	}, ctx)
	return rs
}
