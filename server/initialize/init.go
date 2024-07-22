package initialize

import (
	"go.uber.org/zap"

	"frp-platform/global"
)

func Init() {
	global.VP = initViper()
	global.LOG = initZap()
	zap.ReplaceGlobals(global.LOG)
	global.DB = initGorm()
	if global.DB != nil {
		registerTables()
		// global.LOG.Error("initialize gorm mysql error")
		// os.Exit(0)
	}

	// TODO 初始化redis、mongo

}
