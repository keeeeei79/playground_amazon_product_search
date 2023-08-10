package main

import (
	"fmt"
	"net"
	"os"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/keeeeei79/playground_amazon_product_search/client"
	"github.com/keeeeei79/playground_amazon_product_search/index"
	"github.com/keeeeei79/playground_amazon_product_search/logging"
	pb "github.com/keeeeei79/playground_amazon_product_search/proto"
	"github.com/keeeeei79/playground_amazon_product_search/service"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
				esCli, err := client.NewESIndexClient(cfg, indexName)
				if err != nil {
					logging.Logger.Error("Fail to NewESClient", zap.Error(err))
					return cli.NewExitError(err, 1)
				}

				err = index.BuildIndex(c.Context, esCli)
				if err != nil {
					logging.Logger.Error("Fail to build index", zap.Error(err))
					return cli.NewExitError(err, 1)
				}
				return nil
			},
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "start search server",
			Action: func(c *cli.Context) error {
				listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
				if err != nil {
					logging.Logger.Error("Fail to listen", zap.Error(err))
					return cli.NewExitError(err, 1)
				}
			
				// initialize
				cfg := elasticsearch.Config{
					Addresses: []string{esAddress},
				}
				esCli, err := client.NewESSearchClient(cfg, indexName)
				if err != nil {
					logging.Logger.Error("Fail to NewESClient", zap.Error(err))
					return cli.NewExitError(err, 1)
				}
				s := service.NewSearchService(esCli)
			
				server := grpc.NewServer()
				pb.RegisterSearchServiceServer(server, s)
				reflection.Register(server)
				logging.Logger.Info("Start listening ....")
				server.Serve(listenPort)
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}