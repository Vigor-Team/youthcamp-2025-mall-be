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

package repository

import (
	po2 "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/po"
	"os"

	categoryrepo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/category/repository"
	productrepo "github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/domain/product/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	initDB()
	register()
}

func register() {
	productrepo.GetFactory().SetStockRepository(&StockRepositoryImpl{
		db: DB,
	})
	productrepo.GetFactory().SetProductRepository(&ProductRepositoryImpl{
		db: DB,
	})
	categoryrepo.GetFactory().SetCategoryRepository(&CategoryRepositoryImpl{
		db: DB,
	})
}

func initDB() {
	//dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))
	dsn := "root:root@tcp(127.0.0.1:3306)/product?charset=utf8mb4&parseTime=True&loc=Local"
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
		//nolint:errcheck
		if !DB.Migrator().HasTable(&po2.Category{}) {
			DB.AutoMigrate(
				&po2.Product{},
				&po2.Category{},
			)
			DB.CreateInBatches([]po2.Category{
				{
					Name:        "电子产品",
					Description: "各种电子设备，包括手机、电脑、电视等",
				},
				{
					Name:        "家居用品",
					Description: "提供舒适家居生活所需的各类家具与用品",
				},
				{
					Name:        "食品饮料",
					Description: "各类日常食品与饮料，保证新鲜与健康",
				},
				{
					Name:        "服装鞋帽",
					Description: "时尚服饰与鞋类，适合各种场合穿着",
				},
				{
					Name:        "运动器材",
					Description: "专业运动器材，满足运动需求",
				},
				{
					Name:        "图书音像",
					Description: "各类书籍、音乐与电影，丰富你的精神生活",
				},
			}, 6)
		}

	}

	//if err := DB.Use(tracing.NewPlugin(tracing.WithoutMetrics(), tracing.WithTracerProvider(mtl.TracerProvider))); err != nil {
	//	panic(err)
	//}
}
