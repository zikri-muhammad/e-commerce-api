package configs

import (
	"log"
	"os"
	"strconv"
	"time"
)

// DB holds the DB configuration
type DB struct {
	Host            string
	Port            int
	SslMode         string
	Name            string
	User            string
	Password        string
	Debug           bool
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifetime time.Duration
}

var db = &DB{}

// DBCfg returns the default DB configuration
func DBCfg() *DB {
	return db
}

// LoadDBCfg loads DB configuration
func LoadDBCfg() {
	db.Host = os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Printf("Warning: Invalid DB_PORT, using default 3306: %v", err)
		port = 3306
	}
	db.Port = port

	// Log loaded configuration
	log.Printf("Loading DB Configuration:")
	log.Printf("DB_HOST: %s", db.Host)
	log.Printf("DB_PORT: %d", db.Port)
	log.Printf("DB_USER: %s", os.Getenv("DB_USER"))
	log.Printf("DB_NAME: %s", os.Getenv("DB_NAME"))

	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.Name = os.Getenv("DB_NAME")
	db.SslMode = os.Getenv("DB_SSL_MODE")

	db.Debug, _ = strconv.ParseBool(os.Getenv("DB_DEBUG"))

	maxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	if err != nil {
		log.Printf("Warning: Invalid DB_MAX_OPEN_CONNECTIONS, using default 100: %v", err)
		maxOpen = 100
	}
	db.MaxOpenConn = maxOpen

	maxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS"))
	if err != nil {
		log.Printf("Warning: Invalid DB_MAX_IDLE_CONNECTIONS, using default 10: %v", err)
		maxIdle = 10
	}
	db.MaxIdleConn = maxIdle

	lifeTime, err := strconv.Atoi(os.Getenv("DB_MAX_LIFETIME_CONNECTIONS"))
	if err != nil {
		log.Printf("Warning: Invalid DB_MAX_LIFETIME_CONNECTIONS, using default 3600: %v", err)
		lifeTime = 3600
	}
	db.MaxConnLifetime = time.Duration(lifeTime) * time.Second
}
