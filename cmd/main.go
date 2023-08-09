package main

import (
	"os"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/keeeeei79/playground_amazon_product_search/index"
	"github.com/keeeeei79/playground_amazon_product_search/logging"
	"github.com/keeeeei79/playground_amazon_product_search/search"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	Version = "0.0.1"
)

const (
	name  = "playground"
	usage = ""
	port      = 50051
	esAddress = "http://localhost:9200"
	indexName = "products"
)

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Name = name
	app.Usage = usage
	app.Commands = []*cli.Command{
		{
			Name:    "index",
			Aliases: []string{"i"},
			Usage:   "build index",
			Action: func(c *cli.Context) error {
				logging.Logger.Info("Starting to build index...")
				cfg := elasticsearch.Config{
					Addresses: []string{esAddress},
				}
				esCli, err := search.NewESClient(cfg, indexName)
				if err != nil {
					logging.Logger.Error("Fail to NewESClient", zap.Error(err))
					return cli.NewExitError(err, 1)
				}

				err = index.BuildIndex(c, esCli)
				if err != nil {
					logging.Logger.Error("Fail to build index", zap.Error(err))
					return cli.NewExitError(err, 1)
				}
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}