package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDb interface {
	Init() (*gorm.DB, error)
}

type Db struct {
	Host     string
	User     string
	Port     string
	Password string
	Name     string
}

type DbPostgresSQL struct {
	Db
	SslMode  string
	TimeZone string
}

type DbMySQL struct {
	Db
	Charset   string
	ParseTime string
	Loc       string
}

func (c *DbPostgresSQL) Init() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", c.Host, c.User, c.Password, c.Name, c.Port, c.SslMode, c.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *DbMySQL) Init() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s", c.User, c.Password, c.Host, c.Port, c.Name, c.Charset, c.ParseTime, c.Loc)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
