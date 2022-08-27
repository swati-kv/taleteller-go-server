package app

import (
	"fmt"
	"github.com/spf13/viper"
)

type serviceConfig struct {
	Sample      string
	Environment string
	user        string
	password    string
	localhost   string
	port        int
	dbName      string
}

func InitServiceConfig() serviceConfig {
	return serviceConfig{
		Sample:      ReadEnvString("SAMPLE"),
		Environment: ReadEnvString("ENVIRONMENT"),
		user:        ReadEnvString("DB_USER"),
		password:    ReadEnvString("DB_PASSWORD"),
		localhost:   ReadEnvString("DB_HOST"),
		port:        ReadEnvInt("DB_PORT"),
		dbName:      ReadEnvString("DB_NAME"),
	}
}

type ServiceConfig interface {
}

func (s *serviceConfig) GetSample() string {
	return s.Sample
}

func (s *serviceConfig) GetEnv() string {
	return s.Environment
}

func (s *serviceConfig) GetUser() string {
	return s.user
}

func (s *serviceConfig) GetPassword() string {
	return s.password
}

func (s *serviceConfig) GetHost() string {
	return s.localhost
}
func (s *serviceConfig) GetPort() int {
	return s.port
}

func (s *serviceConfig) GetDbName() string {
	return s.dbName
}

func ReadEnvString(key string) string {
	CheckIfSet(key)
	return viper.GetString(key)
}
func ReadEnvInt(key string) int {
	CheckIfSet(key)
	return viper.GetInt(key)
}

func CheckIfSet(key string) {
	if !viper.IsSet(key) {
		err := fmt.Errorf("key %s is not set", key)
		panic(err)
	}
}
