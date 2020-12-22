package config

import (
	"fmt"
	"log"

	"github.com/joeshaw/envdecode"
)

// Config has DB informations for connection
type Config struct {
	DbUser     string `env:"POSTGRES_USER,required"`
	DbPassword string `env:"POSTGRES_PASSWORD,required"`
	DbHost     string `env:"POSTGRES_HOST,default=localhost"`
	DbPort     uint16 `env:"POSTGRES_PORT,default=5432"`
	DbName     string `env:"POSTGRES_DB,required"`
}

// Get return App configuration
func Get() *Config {
	var conf Config

	if err := envdecode.Decode(&conf); err != nil {
		log.Fatalln(err)
	}

	return &conf
}

// GetDBConnStr return DB information for connection
func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.DbHost, c.DbName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.DbUser,
		c.DbPassword,
		dbhost,
		c.DbPort,
		dbname,
	)
}
