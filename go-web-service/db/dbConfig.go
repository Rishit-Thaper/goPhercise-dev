package db

import (
	env "go-web-service/utils"
)

type DBConfig struct {
	Username string
	Password string
	DBUrl    string
	DBName   string
}

// access the db credentials, like we do in node/express code
func GetDBConfig() DBConfig {
	return DBConfig{
		Username: env.GetEnvVariables("DB_USERNAME"),
		Password: env.GetEnvVariables("DB_PASSWORD"),
		DBUrl:    env.GetEnvVariables("DB_URL"),
		DBName:   env.GetEnvVariables("DB_NAME"),
	}
}
