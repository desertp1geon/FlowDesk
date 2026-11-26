package config

import (
	"fmt"
	"os"

	"github.com/desertp1geon/FlowDesk/pkg/results"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	DbConnectionString string
}

type Config struct {
	DbConfing DbConfig
}

func getVar(key string) results.Result[string] {
	value := os.Getenv(key)
	if value == "" {
		return results.Error[string](fmt.Errorf("Key %s not found", key))
	}
	return results.Ok(value)
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		panic("Cannot load .env")
	}
	dbConStr := getVar("DB_CONNECTION_STRING")
	if dbConStr.IsErr() {
		panic(dbConStr.Err)
	}
	return Config{
		DbConfing: DbConfig{
			DbConnectionString: dbConStr.Value,
		},
	}
}
