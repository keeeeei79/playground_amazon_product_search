package client

import (
	"bytes"
	"context"
	"encoding/json"
	"text/template"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/keeeeei79/playground_amazon_product_search/model"
	"github.com/pkg/errors"
)

const qtmplPath = "client/template/search_dsl.tmpl"

type SearchClient interface {
	Search(context.Context, *model.Query) ([]*model.Doc, error)
}

type ESSearchClient struct {
	indexName string
	cli    *elasticsearch.Client
	qtmpl *template.Template
	cvtr *Converter
}

func NewESSearchClient(cfg elasticsearch.Config, indexName string) (SearchClient, error) {
	cli, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cvtr := NewConverter()
	qtmpl := template.Must(template.ParseFiles(qtmplPath))
	return &ESSearchClient{indexName: indexName, cli: cli, qtmpl: qtmpl, cvtr:cvtr}, nil
}

func (c *ESSearchClient) Search(ctx context.Context, query *model.Query) ([]*model.Doc, error) {
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

func (c *ESSearchClient) buildQuery(query *model.Query) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := c.qtmpl.Execute(&buf, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &buf, nil
}