package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string
	ServiceName string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	ClickHouseHost     string
	ClickHousePort     int
	ClickHouseUser     string
	ClickHousePassword string
	ClickHouseDatabase string

	LogLevel string

	KafkaUrl string

	HttpHost string
	HttpPort string
}

func Load() Config {
	envFileName := cast.ToString(getOrReturnDefault("ENV_FILE_PATH", "./app/.env"))

	if err := godotenv.Load(envFileName); err != nil {
		fmt.Println("No .env file found")
	}
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "info"))
	config.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", ""))

	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DB", "postgres"))

	config.ClickHouseHost = cast.ToString(getOrReturnDefault("CLICKHOUSE_HOST", "localhost"))
	config.ClickHousePort = cast.ToInt(getOrReturnDefault("CLICKHOUSE_PORT", 8123))
	config.ClickHouseUser = cast.ToString(getOrReturnDefault("CLICKHOUSE_USER", "ClickHouse"))
	config.ClickHousePassword = cast.ToString(getOrReturnDefault("CLICKHOUSE_PASSWORD", "ClickHouse"))
	config.ClickHouseDatabase = cast.ToString(getOrReturnDefault("CLICKHOUSE_DATABASE", "ClickHouse"))

	config.KafkaUrl = cast.ToString(getOrReturnDefault("KAFKA_URL", "localhost:9092"))

	config.HttpPort = cast.ToString(getOrReturnDefault("GRPC_PORT", ":8017"))
	config.HttpHost = cast.ToString(getOrReturnDefault("LISTEN_HOST", "localhost"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
