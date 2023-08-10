package client

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/pkg/errors"

	"github.com/keeeeei79/playground_amazon_product_search/model"
)
type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

type ESResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				ID   string    `json:"id"`
				Title string    `json:"title"`
				Description string    `json:"description"`
				BulletPoint string    `json:"bullet_point"`
				Brand string    `json:"brand"`
				Color string    `json:"color"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (cvtr *Converter) convertToDocs(res *esapi.Response) ([]*model.Doc, error) {
	var r ESResponse
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, errors.WithStack(err)
	}
	documents := []*model.Doc{}
	for _, hit := range r.Hits.Hits {
		documents = append(documents, &model.Doc{
			ID:       hit.Source.ID,
			Title:    hit.Source.Title,
			Description:     hit.Source.Description,
			BulletPoint:     hit.Source.BulletPoint,
			Brand:     hit.Source.Brand,
			Color:     hit.Source.Color,
		})
	}
	return documents, nil
}
