package configs

import (
	"fmt"
	"strconv"
)

type DatabaseConfig struct {
	Hostname     string `json:"MONGO_HOSTNAME"`
	Port         int    `json:"MONGO_PORT"`
	DatabaseName string `json:"MONGO_DB_NAME"`
	Username     string `json:"MONGO_DB_USER"`
	Password     string `json:"MONGO_DB_PASSWORD"`
}

func GetDatabaseConfig() DatabaseConfig {
	host := Get("MONGO_HOSTNAME", "db")
	port := Get("MONGO_PORT", "27017")
	dbName := Get("MONGO_DB_NAME", "mongo_db_name")
	username := Get("MONGO_DB_USER", "mongo_username")
	password := Get("MONGO_DB_PASSWORD", "mongo_password")
	portNum, err := strconv.Atoi(port)

	if err != nil {
		panic(fmt.Sprintf("invalid database port config: %s", port))
	}

	return DatabaseConfig{
		Hostname:     host,
		Port:         portNum,
		DatabaseName: dbName,
		Username:     username,
		Password:     password,
	}
}

func (dbConfig DatabaseConfig) GetURI() string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Hostname,
		dbConfig.DatabaseName)
}

func (dbConfig DatabaseConfig) GetURILocal() string {
	return "mongodb://localhost:27017/local_db"
}
