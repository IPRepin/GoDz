package db

import (
	"godz/6-order-api-cart/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDB(conf configs.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.DbConf.DNS), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
