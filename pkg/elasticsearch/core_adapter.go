package elasticsearch

import (
	"context"
	"crypto/tls"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"grabber-match/internal/dtos"
	"grabber-match/pkg/cfg"
	"net/http"
)

const (
	OpTypeIndex = "index"
	PostIndex   = "post_index_v4"
	VideIndex   = "fa_class_video"
)

type CoreElkClient interface {
	Search(ctx context.Context, indexNameOrAlias string, query elastic.Query, pretty bool) (*elastic.SearchResult, error)
	SearchPagination(ctx context.Context, indexNameOrAlias string, query elastic.Query, pagingParams dtos.PaginationFilter, pretty bool) (*elastic.SearchResult, error)
	CheckIndexExist(ctx context.Context, indexName string) (bool, error)
	CreateIndex(ctx context.Context, indexName string, alias *string) error
	CreateIndexIfNotExist(ctx context.Context, indexName string, alias *string) error
	DeleteIndex(ctx context.Context, indexName string) error
	DeleteIndexIfExist(ctx context.Context, indexName string) error
	BulkSaveToElasticsearch(ctx context.Context, indexName string, data map[string]interface{}) error
	BulkUpdateToElasticsearch(ctx context.Context, indexName string, data map[string]interface{}) error
	BulkUpsertToElasticsearch(ctx context.Context, indexName string, data map[string]interface{}) error
	BulkDeleteToElasticsearch(ctx context.Context, indexName string, ids []string) error
	SaveToElasticsearch(ctx context.Context, indexName string, data interface{}) error
	UpdateToElasticsearch(ctx context.Context, indexName string, id string, data interface{}) error
	UpsertToElasticsearch(ctx context.Context, indexName string, id string, data interface{}) error
	DeleteFromElasticsearch(ctx context.Context, indexName string, id string) error
}

type coreElkClient struct {
	client *elastic.Client
}

func NewCoreElkClient(cfgReader *viper.Viper) (CoreElkClient, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}
	url := cfgReader.GetString(cfg.ElasticsearchUrl)
	userName := cfgReader.GetString(cfg.ElasticsearchUserName)
	password := cfgReader.GetString(cfg.ElasticsearchPassword)
	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetHttpClient(httpClient),
		elastic.SetBasicAuth(userName, password),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	if err != nil {
		return nil, err
	}

	return &coreElkClient{
		client: client,
	}, nil
}

func (_this *coreElkClient) Search(ctx context.Context, indexNameOrAlias string, query elastic.Query, pretty bool) (*elastic.SearchResult, error) {
	return _this.client.Search().Index(indexNameOrAlias).Query(query).Pretty(pretty).Do(ctx)
}

func (_this *coreElkClient) SearchPagination(ctx context.Context, indexNameOrAlias string, query elastic.Query, pagingParams dtos.PaginationFilter, pretty bool) (*elastic.SearchResult, error) {
	return _this.client.Search().Index(indexNameOrAlias).Query(query).From(int(pagingParams.PageOffset)).Size(int(pagingParams.PageSize)).Pretty(pretty).Do(ctx)
}

func (_this *coreElkClient) CheckIndexExist(ctx context.Context, indexName string) (bool, error) {
	return elastic.NewIndicesExistsService(_this.client).Index([]string{indexName}).Do(ctx)
}

func (_this *coreElkClient) CreateIndex(ctx context.Context, indexName string, alias *string) error {
	_, err := _this.client.CreateIndex(indexName).Do(ctx)
	if err != nil {
		return err
	}
	if alias != nil && *alias != "" {
		_, err := _this.client.Alias().Add(indexName, *alias).Do(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (_this *coreElkClient) CreateIndexIfNotExist(ctx context.Context, indexName string, alias *string) error {
	exists, err := _this.CheckIndexExist(ctx, indexName)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	return _this.CreateIndex(ctx, indexName, alias)
}

func (_this *coreElkClient) DeleteIndex(ctx context.Context, indexName string) error {
	_, err := _this.client.DeleteIndex(indexName).Do(ctx)
	return err
}

func (_this *coreElkClient) DeleteIndexIfExist(ctx context.Context, indexName string) error {
	exists, err := _this.CheckIndexExist(ctx, indexName)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	return _this.DeleteIndex(ctx, indexName)
}

func (_this *coreElkClient) BulkSaveToElasticsearch(ctx context.Context, indexName string, data map[string]interface{}) error {
	bulk := _this.client.Bulk()
	for id, object := range data {
		bulk = bulk.Add(elastic.NewBulkIndexRequest().OpType(OpTypeIndex).Index(indexName).Id(id).Doc(object))
	}
	_, err := bulk.Do(ctx)
	if err != nil {
		return err
	}
	_, err = _this.client.Flush().Index(indexName).Do(ctx)
	return err
}

func (_this *coreElkClient) BulkUpdateToElasticsearch(ctx context.Context, indexName string, data map[string]interface{}) error {
	bulk := _this.client.Bulk()
	for id, object := range data {
		bulk = bulk.Add(elastic.NewBulkUpdateRequest().Index(indexName).Id(id).Doc(object))
	}
	_, err := bulk.Do(ctx)
	if err != nil {
		return err
	}
	_, err = _this.client.Flush().Index(indexName).Do(ctx)
	return err
}

func (_this *coreElkClient) BulkUpsertToElasticsearch(ctx context.Context, indexName string, data map[string]interface{}) error {
	bulk := _this.client.Bulk()
	for id, object := range data {
		bulk = bulk.Add(elastic.NewBulkUpdateRequest().Index(indexName).Id(id).Doc(object).DocAsUpsert(true))
	}
	_, err := bulk.Do(ctx)
	if err != nil {
		return err
	}
	_, err = _this.client.Flush().Index(indexName).Do(ctx)
	return err
}

func (_this *coreElkClient) BulkDeleteToElasticsearch(ctx context.Context, indexName string, ids []string) error {
	bulk := _this.client.Bulk()
	for _, id := range ids {
		bulk = bulk.Add(elastic.NewBulkDeleteRequest().Index(indexName).Id(id))
	}
	_, err := bulk.Do(ctx)
	if err != nil {
		return err
	}
	_, err = _this.client.Flush().Index(indexName).Do(ctx)
	return err
}

func (_this *coreElkClient) SaveToElasticsearch(ctx context.Context, indexName string, data interface{}) error {
	_, err := _this.client.Index().Index(indexName).BodyJson(data).Do(ctx)
	if err != nil {
		return err
	}
	_, err = _this.client.Flush().Index(indexName).Do(ctx)
	return err
}

func (_this *coreElkClient) UpdateToElasticsearch(ctx context.Context, indexName string, id string, data interface{}) error {
	_, err := _this.client.Update().Index(indexName).Id(id).Doc(data).Do(ctx)
	if err != nil {
		return err
	}
	_, err = _this.client.Flush().Index(indexName).Do(ctx)
	return err
}

func (_this *coreElkClient) UpsertToElasticsearch(ctx context.Context, indexName string, id string, data interface{}) error {
	_, err := _this.client.Update().Index(indexName).Id(id).Doc(data).DocAsUpsert(true).Do(ctx)
	if err != nil {
		return err
	}
	_, err = _this.client.Flush().Index(indexName).Do(ctx)
	return err
}

func (_this *coreElkClient) DeleteFromElasticsearch(ctx context.Context, indexName string, id string) error {
	_, err := _this.client.Delete().Index(indexName).Id(id).Do(ctx)
	if err != nil {
		return err
	}
	_, err = _this.client.Flush().Index(indexName).Do(ctx)
	return err
}
