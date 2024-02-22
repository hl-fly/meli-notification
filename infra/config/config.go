package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Read() *Config {
	godotenv.Load(".env")

	var config Config
	config.Cache.Port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
	config.Cache.Host = os.Getenv("REDIS_HOSTNAME")
	config.Cache.Prefix = os.Getenv("REDIS_PREFIX")
	config.Cache.Password = os.Getenv("REDIS_PASSWORD")

	config.DB.User = os.Getenv("DATABASE_USER")
	config.DB.Password = os.Getenv("DATABASE_PASSWORD")
	config.DB.Address = os.Getenv("DATABASE_ADDRESS")
	config.DB.Port = os.Getenv("DATABASE_PORT")
	config.DB.Name = os.Getenv("DATABASE_NAME")

	config.App.Debug = os.Getenv("ENV_DEBUG")

	config.Server.WWWSources = os.Getenv("WWW_SOURCES")
	config.Server.Port = os.Getenv("PORT")
	config.Server.Auth.JWTSignatureKey = os.Getenv("JWT_SIGNATURE_KEY")

	return &config
}
