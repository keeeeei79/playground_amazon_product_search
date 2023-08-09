package index

import (
	"encoding/csv"
	"io"
	"os"
	"time"

	"github.com/keeeeei79/playground_amazon_product_search/logging"
	"github.com/keeeeei79/playground_amazon_product_search/model"
	"github.com/keeeeei79/playground_amazon_product_search/search"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const csvFile = "data/products_train.csv"


func BuildIndex(c *cli.Context, searchCli search.Client) error {
	start := time.Now()
	docs, err := readCSV(c)
	if err != nil {
		return errors.WithStack(err)
	}
	printDocs(docs)
	err = searchCli.Index(docs)
	if err != nil {
		return errors.WithStack(err)
	}
	elapsed := time.Since(start)
	logging.Logger.Info("Took time", zap.Duration("elapsed", elapsed))
	return nil

}

func readCSV(c *cli.Context) ([]*model.Doc, error) {
	file, err := os.Open(csvFile)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer file.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(file)

	// ヘッダー行をスキップ
	_, err = reader.Read()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// レコードを一行ずつ読み込む
	var docs []*model.Doc
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				// 全て読み終われば抜ける
				break
			}
			if err == csv.ErrFieldCount {
				logging.Logger.Error("Skip malformed line", zap.Error(err))
				continue
			}

			if err != nil {
				return nil, errors.WithStack(err)
			}
		}

		docs = append(docs, &model.Doc{
			ID:          record[0],
			Title:       record[1],
			Description: record[2],
			BulletPoint: record[3],
			Brand:       record[4],
			Color:       record[5],
		}) 
	}
	logging.Logger.Info("CSV File size", zap.Int("Row", len(docs)))
	return docs, nil
}


func printDocs(docs []*model.Doc){
	for i, doc := range docs {
		logging.Logger.Info("Doc", zap.String("ID",doc.ID), zap.String("Title", doc.Title))
		if (i == 10) {
			break
		}
	}
}
