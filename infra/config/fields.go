package config

type Config struct {
	App    AppConfig
	Server ServerConfig
	DB     DBConfig
	Cache  CacheConfig
}

type AppConfig struct {
	Debug string
}

type ServerConfig struct {
	Port       string
	WWWSources string
	Auth       ServerAuthConfig
}

type ServerAuthConfig struct {
	JWTSignatureKey string
}

type DBConfig struct {
	User     string
	Password string
	Address  string
	Port     string
	Name     string
}

type CacheConfig struct {
	Host     string
	Port     int
	Prefix   string
	Password string
}
