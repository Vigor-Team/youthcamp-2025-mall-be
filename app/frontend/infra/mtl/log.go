// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mtl

import (
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/frontend/conf"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/server"
	hertzzap "github.com/hertz-contrib/logger/zap"
	hertobszap "github.com/hertz-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initLog() {
	var opts []hertzzap.Option
	var output zapcore.WriteSyncer
	if os.Getenv("GO_ENV") != "online" {
		opts = append(opts, hertzzap.WithCoreEnc(zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())))
		output = &zapcore.BufferedWriteSyncer{
			WS: zapcore.AddSync(&lumberjack.Logger{
				Filename:   conf.GetConf().Hertz.LogFileName,
				MaxSize:    conf.GetConf().Hertz.LogMaxSize,
				MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
				MaxAge:     conf.GetConf().Hertz.LogMaxAge,
			}),
		}
	} else {
		opts = append(opts, hertzzap.WithCoreEnc(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())))
		// async log
		output = &zapcore.BufferedWriteSyncer{
			WS:            zapcore.AddSync(os.Stdout),
			FlushInterval: time.Minute,
		}
		server.RegisterShutdownHook(func() {
			output.Sync() //nolint:errcheck
		})
	}
	log := hertobszap.NewLogger(hertobszap.WithLogger(hertzzap.NewLogger(opts...)))
	hlog.SetLogger(log)
	hlog.SetLevel(hlog.LevelInfo)
	hlog.SetOutput(output)
}
