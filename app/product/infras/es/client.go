package es

import (
	"context"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/conf"
	"github.com/elastic/go-elasticsearch/v8"
	"os"
	"sync"
)

var (
	esOnce sync.Once
	esCli  *elasticsearch.TypedClient
)

func GetESClient() *elasticsearch.TypedClient {
	if esCli != nil {
		return esCli
	}
	esOnce.Do(func() {
		cfg := elasticsearch.Config{
			Addresses: []string{conf.GetConf().ES.Address},
			Username:  conf.GetConf().ES.UserName,
			Password:  os.Getenv("ES_PASSWORD"),
		}
		cli, err := elasticsearch.NewTypedClient(cfg)
		if err != nil {
			panic("new es client failed, err=" + err.Error())
		}
		_, err = cli.Info().Do(context.Background())
		if err != nil {
			panic("client connect fail, err=" + err.Error())
		}
		esCli = cli
	})
	return esCli
}
