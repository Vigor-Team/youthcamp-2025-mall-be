package es

import (
	"context"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/common/model/entity"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/embedding"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/infras/repository/converter"
	"github.com/bytedance/sonic"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"strconv"
)

type ProductESClient struct{}

var productESClient ProductESClient

func GetProductESClient() *ProductESClient {
	return &productESClient
}

func (c *ProductESClient) UpsertProduct(ctx context.Context, productId uint32, product *entity.ProductES) error {
	doc := getDocFromProductES(product)
	_, err := GetESClient().Update("product", strconv.FormatInt(int64(productId), 10)).Doc(doc).Upsert(doc).Refresh(refresh.Refresh{Name: "true"}).Do(ctx)
	return err
}

func (c *ProductESClient) BatchGetProductById(ctx context.Context, productIds []uint32) ([]*entity.ProductEntity, error) {
	termQuery := map[string]types.TermsQueryField{}
	ids := make([]interface{}, len(productIds))
	for i, id := range productIds {
		ids[i] = strconv.FormatInt(int64(id), 10)
	}
	termQuery["_id"] = ids
	resp, err := GetESClient().Search().
		Index("product").
		Request(&search.Request{
			Query: &types.Query{
				Terms: &types.TermsQuery{
					TermsQuery: termQuery,
				},
			},
		}).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("search request error: %w", err)
	}
	products := make([]*entity.ProductEntity, 0)
	for _, hit := range resp.Hits.Hits {
		do := converter.ProductDoWithESConverter.Convert2DO(ctx, getProductESFormSource(string(hit.Source_)))
		products = append(products, do)
	}
	return products, nil
}

func (c *ProductESClient) SearchProduct(ctx context.Context, keyword string) ([]*entity.ProductEntity, error) {
	eb, err := embedding.GetArkEmbedding(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("create embedding error: %w", err)
	}
	vectors, err := eb.EmbedStrings(ctx, []string{keyword})
	if err != nil {
		return nil, fmt.Errorf("embedding error: %w", err)
	}
	if len(vectors) == 0 || len(vectors[0]) == 0 {
		return nil, fmt.Errorf("empty embedding result")
	}
	queryVector := make([]float32, len(vectors[0]))
	for i, v := range vectors[0] {
		queryVector[i] = float32(v)
	}
	size := 100
	k := 50
	numCandidates := 100
	var boost float32 = 0.5
	resp, err := GetESClient().Search().
		Index("product").
		Request(&search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Should: []types.Query{
						{
							MultiMatch: &types.MultiMatchQuery{
								Query:  keyword,
								Fields: []string{"name^3", "description^2", "spuName"},
							},
						},
						{
							Knn: &types.KnnQuery{
								Field:         "embedding",
								QueryVector:   queryVector,
								K:             &k,
								NumCandidates: &numCandidates,
								Boost:         &boost,
							},
						},
					},
				},
			},
			Size: &size,
		}).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("search request error: %w", err)
	}

	products := make([]*entity.ProductEntity, 0, len(resp.Hits.Hits))
	for _, hit := range resp.Hits.Hits {
		do := converter.ProductDoWithESConverter.Convert2DO(ctx, getProductESFormSource(string(hit.Source_)))
		products = append(products, do)
	}
	return products, nil
}

func getProductESFormSource(source string) *entity.ProductES {
	sourceMap := make(map[string]interface{})
	_ = sonic.UnmarshalString(source, &sourceMap)
	categoryNames := sourceMap["category_names"].([]interface{})
	categoryNamesStr := make([]string, len(categoryNames))
	for i, v := range categoryNames {
		categoryNamesStr[i] = v.(string)
	}
	ret := &entity.ProductES{
		ID:            uint32(sourceMap["id"].(float64)),
		Name:          sourceMap["name"].(string),
		Description:   sourceMap["description"].(string),
		Picture:       sourceMap["picture"].(string),
		Price:         float32(sourceMap["price"].(float64)),
		Stock:         uint32(sourceMap["stock"].(float64)),
		SpuName:       sourceMap["spu_name"].(string),
		SpuPrice:      float32(sourceMap["spu_price"].(float64)),
		Status:        uint32(sourceMap["status"].(float64)),
		CategoryNames: categoryNamesStr,
	}
	return ret
}

func getDocFromProductES(entity *entity.ProductES) map[string]interface{} {
	doc := make(map[string]interface{})
	doc["id"] = entity.ID
	doc["name"] = entity.Name
	doc["description"] = entity.Description
	doc["picture"] = entity.Picture
	doc["price"] = entity.Price
	doc["stock"] = entity.Stock
	doc["spu_name"] = entity.SpuName
	doc["spu_price"] = entity.SpuPrice
	doc["status"] = entity.Status
	doc["category_names"] = entity.CategoryNames
	return doc
}
