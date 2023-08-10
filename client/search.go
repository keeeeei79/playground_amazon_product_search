package client

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/keeeeei79/playground_amazon_product_search/model"
	"github.com/pkg/errors"
)


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