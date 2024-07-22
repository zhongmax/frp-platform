package initialize

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"

	"frp-platform/global"
	"frp-platform/utils"
)

func initZap() (logger *zap.Logger) {
	exist, _ := utils.PathExists(global.CONFIG.Zap.Director)
	if !exist {
		fmt.Printf("create %v directory\n", global.CONFIG.Zap.Director)
		_ = os.Mkdir(global.CONFIG.Zap.Director, os.ModePerm)
	}
	levels := global.CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := newZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	if global.CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

type zapCore struct {
	level zapcore.Level
	zapcore.Core
}

func newZapCore(level zapcore.Level) *zapCore {
	entity := &zapCore{level: level}
	syncer := entity.writeSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	entity.Core = zapcore.NewCore(global.CONFIG.Zap.Encoder(), syncer, levelEnabler)
	return entity
}

func (z *zapCore) writeSyncer(formats ...string) zapcore.WriteSyncer {
	cutter := newCutter(
		global.CONFIG.Zap.Director,
		z.level.String(),
		global.CONFIG.Zap.RetentionDay,
		CutterWithLayout(time.DateOnly),
		CutterWithFormats(formats...),
	)
	if global.CONFIG.Zap.LogInConsole {
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		return zapcore.AddSync(multiSyncer)
	}
	return zapcore.AddSync(cutter)
}

func (z *zapCore) Enabled(level zapcore.Level) bool {
	return z.level == level
}

func (z *zapCore) With(fields []zapcore.Field) zapcore.Core {
	return z.Core.With(fields)
}

func (z *zapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *zapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	for i := 0; i < len(fields); i++ {
		if fields[i].Key == "business" || fields[i].Key == "folder" || fields[i].Key == "directory" {
			syncer := z.writeSyncer(fields[i].String)
			z.Core = zapcore.NewCore(global.CONFIG.Zap.Encoder(), syncer, z.level)
		}
	}
	return z.Core.Write(entry, fields)
}

func (z *zapCore) Sync() error {
	return z.Core.Sync()
}
