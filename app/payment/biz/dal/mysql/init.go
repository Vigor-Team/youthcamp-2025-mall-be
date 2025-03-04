package mysql

import (
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/common/mtl"
	"gorm.io/plugin/opentelemetry/tracing"
	"os"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/payment/biz/model"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/payment/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if os.Getenv("GO_ENV") != "online" {
		DB.AutoMigrate( //nolint:errcheck
			&model.PaymentLog{},
		)
	}
	if err = DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
		panic(err)
	}
}
