package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sync"

	"frp-platform/config"
)

var (
	DB     *gorm.DB
	REDIS  redis.UniversalClient
	MONGO  *qmgo.QmgoClient
	CONFIG config.Server
	VP     *viper.Viper
	LOG    *zap.Logger
	// Timer               = timer.NewTimerTask()
	Concurrency_Control = &singleflight.Group{}
	BlackCache          local_cache.Cache

	lock     sync.RWMutex
	validate = validator.New()
)

func Validate(obj any) error {
	return validate.Struct(obj)
}
