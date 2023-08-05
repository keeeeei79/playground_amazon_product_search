package index

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/keeeeei79/playground_amazon_product_search/logging"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

const csvFile = "data/products_train.csv"

type Doc struct {
	ID string
	Title string
	Description string
	BulletPoint string
	Brand string
	Color string
}

func BuildIndex(c *cli.Context) error {
	// esClient
	docs, err := readCSV(c)
	if err != nil {
		return errors.WithStack(err)
	}
	printDocs(docs)
	// postDoc
	return nil

}

func readCSV(c *cli.Context) ([]*Doc, error) {
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
	var docs []*Doc
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

		docs = append(docs, &Doc{
			ID:          record[0],
			Title:       record[1],
			Description: record[2],
			BulletPoint: record[3],
			Brand:       record[4],
			Color:       record[5],
		}) 
	}
	return docs, nil
}


func printDocs(docs []*Doc){
	for i, doc := range docs {
		logging.Logger.Info("Doc", zap.String("ID",doc.ID), zap.String("Title", doc.Title))
		if (i == 10) {
			break
		}
	}
}
