package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
)

var DbConnections map[string]*gorm.DB

func Init() {
	dbConfigurations := map[string]IDb{
		"POSTGRES": &DbPostgresSQL{
			Db: Db{
				Host:     os.Getenv("DB_HOST_POSTGRES"),
				User:     os.Getenv("DB_USER_POSTGRES"),
				Port:     os.Getenv("DB_PORT_POSTGRES"),
				Password: os.Getenv("DB_PASSWORD_POSTGRES"),
				Name:     os.Getenv("DB_NAME_POSTGRES"),
			},
			SslMode:  os.Getenv("DB_SSLMODE_POSTGRES"),
			TimeZone: os.Getenv("DB_TIMEZONE_POSTGRES"),
		},
	}

	DbConnections = make(map[string]*gorm.DB)
	for key, value := range dbConfigurations {
		db, err := value.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", key))
		}
		DbConnections[key] = db
		fmt.Printf("Successfully connected to database %s", key)
	}
}

func Connection(name string) (*gorm.DB, error) {
	if DbConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("Connection is undefined")
	}
	return DbConnections[strings.ToUpper(name)], nil
}
