package initialize

import (
	"go.uber.org/zap"
	"os"

	"frp-platform/global"
)

func registerTables() {
	db := global.DB
	err := db.AutoMigrate(
	// model.Api{},
	// model.User{},
	)
	if err != nil {
		global.LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.LOG.Info("register table success")
}
