package config

import "github.com/ShatteredRealms/go-common-service/pkg/config"

var (
	Version = "v1.0.0"
)

type ConnectionConfig struct {
	config.BaseConfig `yaml:",inline" mapstructure:",squash"`
	Postgres          config.DBPoolConfig `yaml:"postgres"`
}

func NewConnectionConfig() *ConnectionConfig {
	return &ConnectionConfig{
		BaseConfig: config.BaseConfig{
			Server: config.ServerAddress{
				Host: "localhost",
				Port: "8082",
			},
			Keycloak: config.KeycloakConfig{
				BaseURL:      "localhost:8080",
				Realm:        "default",
				Id:           "274ed62e-127b-44d4-9832-c982f45e91c6",
				ClientId:     "sro-connection-service",
				ClientSecret: "**********",
			},
			Mode:                "local",
			LogLevel:            0,
			OpenTelemtryAddress: "localhost:4317",
		},
		Postgres: config.DBPoolConfig{
			Master: config.DBConfig{
				ServerAddress: config.ServerAddress{},
				Name:          "connection-service",
				Username:      "postgres",
				Password:      "password",
			},
		},
	}
}
