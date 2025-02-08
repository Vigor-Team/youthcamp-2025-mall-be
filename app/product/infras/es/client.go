package es

import (
	"github.com/elastic/go-elasticsearch/v8"
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
			Addresses: []string{"http://localhost:9200"},
			Username:  "elastic",
			Password:  "123456",
		}
		cli, err := elasticsearch.NewTypedClient(cfg)
		if err != nil {
			panic("new es client failed, err=" + err.Error())
		}
		esCli = cli
	})
	return esCli
}
