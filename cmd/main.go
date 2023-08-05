package main

import (
	"os"

	"github.com/keeeeei79/playground_amazon_product_search/index"
	"github.com/keeeeei79/playground_amazon_product_search/logging"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var (
	Version = "0.0.1"
)

const (
	name  = "playground"
	usage = ""
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
				err := index.BuildIndex(c)
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