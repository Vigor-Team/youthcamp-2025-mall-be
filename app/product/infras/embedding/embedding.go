package embedding

import (
	"context"
	"os"
	"sync"

	"github.com/cloudwego/eino-ext/components/embedding/ark"
	"github.com/cloudwego/eino/components/embedding"
)

var (
	EB   embedding.Embedder
	once sync.Once
)

func defaultArkEmbeddingConfig(_ context.Context) (*ark.EmbeddingConfig, error) {
	config := &ark.EmbeddingConfig{
		Model:  os.Getenv("ARK_EMBEDDING_MODEL"),
		APIKey: os.Getenv("ARK_API_KEY"),
	}
	return config, nil
}

func GetArkEmbedding(ctx context.Context, config *ark.EmbeddingConfig) (eb embedding.Embedder, err error) {
	once.Do(func() {
		if config == nil {
			config, err = defaultArkEmbeddingConfig(ctx)
			if err != nil {
				return
			}
		}
		EB, err = ark.NewEmbedder(ctx, config)
		if err != nil {
			return
		}
	})
	return EB, err
}
