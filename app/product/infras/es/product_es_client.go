package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/product/model/entity"
	"github.com/bytedance/sonic"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
	"strconv"
)

type ProductESClient struct{}

var productESClient ProductESClient

func GetProductESClient() *ProductESClient {
	return &productESClient
}

func (c *ProductESClient) UpdateProduct(ctx context.Context, productId uint32, product *entity.ProductES) error {
	doc := getDocFromProductES(product)
	_, err := GetESClient().Update("product", strconv.FormatInt(int64(productId), 10)).Doc(doc).Upsert(doc).Refresh(refresh.Refresh{Name: "true"}).Do(ctx)
	return err
}

func (c *ProductESClient) BatchGetProductById(ctx context.Context, productIds []uint32) ([]*entity.ProductES, error) {
	var buf bytes.Buffer
	mgetBody := make(map[string]interface{})
	docs := make([]map[string]string, len(productIds))
	for i, id := range productIds {
		docs[i] = map[string]string{
			"_index": "product",
			"_id":    strconv.FormatInt(int64(id), 10),
		}
	}
	mgetBody["docs"] = docs
	if err := json.NewEncoder(&buf).Encode(mgetBody); err != nil {
		return nil, fmt.Errorf("encoding error: %w", err)
	}
	req := esapi.MgetRequest{
		Body: &buf,
	}
	res, err := req.Do(ctx, GetESClient())
	if err != nil {
		return nil, fmt.Errorf("mget request error: %w", err)
	}
	defer res.Body.Close()
	var rsp struct {
		Docs []struct {
			Found  bool            `json:"found"`
			Source json.RawMessage `json:"_source"`
		} `json:"docs"`
	}
	if err := json.NewDecoder(res.Body).Decode(&rsp); err != nil {
		return nil, fmt.Errorf("decode error: %w", err)
	}
	entities := make([]*entity.ProductES, 0)
	for _, doc := range rsp.Docs {
		if doc.Found {
			entities = append(entities, getProductESFormSource(string(doc.Source)))
		}
	}
	return entities, nil
}

func (c *ProductESClient) SearchProduct(ctx context.Context, keyword string) ([]*entity.ProductES, error) {
	size := 1000
	resp, err := GetESClient().Search().
		Index("product").
		Request(&search.Request{
			Query: &types.Query{
				MultiMatch: &types.MultiMatchQuery{
					Query:  keyword,
					Fields: []string{"name", "description", "spuName"},
					Type: &textquerytype.TextQueryType{
						Name: "best_fields",
					},
				},
			},
			Size: &size,
		}).
		Do(ctx)
	if err != nil {
		return nil, fmt.Errorf("search request error: %w", err)
	}
	products := make([]*entity.ProductES, 0)
	for _, hit := range resp.Hits.Hits {
		products = append(products, getProductESFormSource(string(hit.Source_)))
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
