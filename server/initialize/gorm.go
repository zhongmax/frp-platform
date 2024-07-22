package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"

	"frp-platform/config"
	"frp-platform/global"
)

func initGorm() *gorm.DB {
	m := global.CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	gormConfig := getGormConfig(m.Prefix, m.Singular)
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig); err != nil {
		panic(fmt.Errorf("gorm open mysql error: %s", err))
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func getGormConfig(prefix string, singular bool) *gorm.Config {
	mysqlConfig := global.CONFIG.Mysql
	return &gorm.Config{
		Logger: logger.New(newGormLogWriter(mysqlConfig, log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      mysqlConfig.LogLevel(),
			Colorful:      true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}

type gormLogWriter struct {
	config config.Mysql
	writer logger.Writer
}

func newGormLogWriter(config config.Mysql, writer logger.Writer) *gormLogWriter {
	return &gormLogWriter{config: config, writer: writer}
}

func (c *gormLogWriter) Printf(message string, data ...any) {
	if c.config.LogZap {
		switch c.config.LogLevel() {
		case logger.Silent:
			zap.L().Debug(fmt.Sprintf(message, data...))
		case logger.Error:
			zap.L().Error(fmt.Sprintf(message, data...))
		case logger.Warn:
			zap.L().Warn(fmt.Sprintf(message, data...))
		case logger.Info:
			zap.L().Info(fmt.Sprintf(message, data...))
		default:
			zap.L().Info(fmt.Sprintf(message, data...))
		}
		return
	}
	c.writer.Printf(message, data...)
}
