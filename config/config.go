package config

import (
	dbPkg "helpstack/config/database"
	"os"
)

func getEnv(key string, defaultVal string) string  {
	if value, exists :=os.LookupEnv(key); exists&& value!=""{
		return value
	}

	return defaultVal
}

type AppConfig struct {
	ServerPort  string
	ServerHost string
	SqlDbConfig dbPkg.SqlDbConfig
}

func New() *AppConfig {
	return &AppConfig{
		ServerPort: getEnv("Server_PORT", "8080"),
		ServerHost: getEnv("Server_Host", ""),
		SqlDbConfig: dbPkg.SqlDbConfig{
			Username: getEnv("PG_USER_NAME","postgres"),
			DbName: getEnv("POSTGRES_DB_NAME","stackhelp"),
			Password: getEnv("POSTGRES_PASSWORD","54321"),
			Host: getEnv("POSTGRES_HOST",""),
			Driver: getEnv("SQL_DRIVER", "postgres"),
		},
	}
}
