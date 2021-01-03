package dbcfg

import (
	"fmt"
	"log"

	"github.com/joeshaw/envdecode"
)

// DBConfig has DB informations for connection
type DBConfig struct {
	DbUser     string `env:"POSTGRES_USER,required"`
	DbPassword string `env:"POSTGRES_PASSWORD,required"`
	DbHost     string `env:"POSTGRES_HOST,default=localhost"`
	DbPort     uint16 `env:"POSTGRES_PORT,default=5432"`
	DbName     string `env:"POSTGRES_DB,required"`
}

// GetDBConfig return App configuration
func GetDBConfig() *DBConfig {
	var conf DBConfig

	if err := envdecode.Decode(&conf); err != nil {
		log.Fatalln(err)
	}

	return &conf
}

// GetDBConnStr return DB information for connection
func (c *DBConfig) GetDBConnStr() string {
	return c.getDBConnStr(c.DbHost, c.DbName)
}

func (c *DBConfig) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.DbUser,
		c.DbPassword,
		dbhost,
		c.DbPort,
		dbname,
	)
}
