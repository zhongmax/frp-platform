package main

import (
	"frp-platform/core"
	"frp-platform/global"
	"frp-platform/initialize"
)

func main() {
	initialize.Init()
	if global.DB != nil {
		db, _ := global.DB.DB()
		defer func() { _ = db.Close() }()
	}
	core.RunServer()
}
