package model

import "time"

type ServiceConfig struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	ServicePort int    `mapstructure:"SERVICE_PORT"`

	DBHost         string        `mapstructure:"MG_HOST"`
	DBPort         int           `mapstructure:"MG_PORT"`
	DBUser         string        `mapstructure:"MG_USER"`
	DBPwd          string        `mapstructure:"MG_PWD"`
	DBName         string        `mapstructure:"MG_DBNAME"`
	AuthSource     string        `mapstructure:"MG_AUTH_SOURCE"`
	ReplicaSet     string        `mapstructure:"MG_REPLICA_SET"`
	MaxPoolSize    int           `mapstructure:"MG_MAX_POOL_SIZE"`
	MinPoolSize    int           `mapstructure:"MG_MIN_POOL_SIZE"`
	ConnectTimeout time.Duration `mapstructure:"MG_CONNECT_TIMEOUT"`
	SocketTimeout  time.Duration `mapstructure:"MG_SOCKET_TIMEOUT"`

	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}
