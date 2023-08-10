package client

import (
	"context"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/keeeeei79/playground_amazon_product_search/model"
	"github.com/pkg/errors"
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
