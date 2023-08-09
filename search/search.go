package search

import (
	"bytes"
	"context"
	"encoding/json"
	"sync/atomic"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"github.com/keeeeei79/playground_amazon_product_search/logging"
	"github.com/keeeeei79/playground_amazon_product_search/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Client interface {
	Index(context.Context, []*model.Doc) error
	Search(context.Context, *model.Query) ([]*model.Doc, error)
}

type ESClient struct {
	cli    *elasticsearch.Client
	indexName string
	cvtr *Converter
}

func NewESClient(cfg elasticsearch.Config, indexName string) (Client, error) {
	cli, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cvtr := NewConverter()
	return &ESClient{cli: cli, indexName: indexName, cvtr: cvtr}, nil
}

func (c *ESClient) Index(ctx context.Context, docs []*model.Doc) error {
	bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index: c.indexName,
		Client: c.cli,
	})
	if err != nil {
		return errors.WithStack(err)
	}
	var countSuccessful uint64
	var countFail uint64
	for i, doc := range docs {
		data, err := json.Marshal(doc)
		if err != nil {
			logging.Logger.Error("Fail to encode to json", zap.String("Doc ID", doc.ID))
			continue
		}
		err = bi.Add(
			ctx,
			esutil.BulkIndexerItem{
				// Action field configures the operation to perform (index, create, delete, update)
				Action: "index",

				// Body is an `io.Reader` with the payload
				Body: bytes.NewReader(data),

				// OnSuccess is called for each successful operation
				OnSuccess: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem) {
					atomic.AddUint64(&countSuccessful, 1)
				},

				// OnFailure is called for each failed operation
				OnFailure: func(ctx context.Context, item esutil.BulkIndexerItem, res esutil.BulkIndexerResponseItem, err error) {
					atomic.AddUint64(&countFail, 1)
					if err != nil {
						logging.Logger.Error("Fail to index", zap.String("Doc ID", doc.ID), zap.Error(err))
					} else {
						logging.Logger.Error("Fail to index", zap.String("Doc ID", doc.ID), zap.String("Error Type",res.Error.Type), zap.String("Error Reason", res.Error.Reason))
					}
				},
			},
		)
		if err != nil {
			logging.Logger.Error("Fail to bulk index", zap.Error(err))
			return errors.WithStack(err)
		}
		if i != 0 && i % 10000 == 0 {
			logging.Logger.Info("Succeed index", zap.Int("i",i))
		}
	}
	logging.Logger.Info("Indexing Result", zap.Uint64("Success", countSuccessful), zap.Uint64("Fail", countFail))
	return nil
}

func (c *ESClient) Search(ctx context.Context, query *model.Query) ([]*model.Doc, error) {
	esQuery, err := c.buildQuery(query)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res, err := c.cli.Search(
		c.cli.Search.WithContext(ctx),
		c.cli.Search.WithIndex(c.indexName),
		c.cli.Search.WithBody(esQuery),
		c.cli.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer res.Body.Close()
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			return nil, errors.WithStack(err)
		}
		return nil, errors.Errorf("[%s] %s: %s",
			res.Status(),
			e["error"].(map[string]interface{})["type"],
			e["error"].(map[string]interface{})["reason"])
	}
	docs, err := c.cvtr.convertToDocs(res)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return docs, nil
}

func (client *ESClient) buildQuery(query *model.Query) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	esQuery := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": query.Keyword,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(esQuery); err != nil {
		return nil, errors.WithStack(err)
	}
	return &buf, nil
}