package main

import (
	"godz/4-order-api/configs"
	"godz/4-order-api/pkg/db"
)

func main() {
	conf := configs.GetConfig()
	_ = db.NewDB(conf)
}
