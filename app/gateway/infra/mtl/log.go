package mtl

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/conf"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/logger/zap"
	hertobszap "github.com/hertz-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func initLog() {
	var opts []hertzzap.Option
	var output zapcore.WriteSyncer
	opts = append(opts, hertzzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
	output = &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
	}
	log := hertobszap.NewLogger(hertobszap.WithLogger(hertzzap.NewLogger(opts...)))
	hlog.SetLogger(log)
	hlog.SetLevel(hlog.LevelInfo)
	hlog.SetOutput(output)
}
