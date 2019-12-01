package server

import "os"

// Config type
type Config struct {
	db     *Database
	server *Server
}

//Database config
type Database struct {
	DBName     string
	DBUser     string
	DBPassword string
}

//Server config
type Server struct {
	BindAddr string
	LogLevel string
}

//NewConfig create a new config instance
func NewConfig() *Config {
	return &Config{
		db: &Database{
			DBName:     getEnv("DB_NAME", "homestage"),
			DBUser:     getEnv("DB_USER", "homestage"),
			DBPassword: getEnv("DB_PASS", "homestage"),
		},
		server: &Server{
			BindAddr: getEnv("BIND_ADDR", "homestage"),
			LogLevel: getEnv("LOG_LEVEL", "homestage"),
		},
	}
}

//Getting enviroment param helper
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}