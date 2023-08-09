package service

import (
	"context"

	"github.com/keeeeei79/playground_amazon_product_search/logging"
	"github.com/keeeeei79/playground_amazon_product_search/model"
	pb "github.com/keeeeei79/playground_amazon_product_search/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Searcher interface {
	Search(ctx context.Context, query *model.Query) ([]*model.Doc, error)
}


type SearchService struct {
	pb.UnimplementedSearchServiceServer
	searcher       Searcher
}

func NewSearchService(searcher Searcher) *SearchService {
	return &SearchService{searcher: searcher}
}

func (s *SearchService) Search(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	query := convertToQuery(req)
	logging.Logger.Info("Request query: ", zap.String("keyword", query.Keyword))

	docs, err := s.searcher.Search(ctx, query)
	if err != nil {
		logging.Logger.Error("Fail to search", zap.Error(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	return convertToPB(docs), nil
}

// adapter
func convertToQuery(req *pb.Request) *model.Query {
	return &model.Query{Keyword: req.Keyword}
}

func convertToPB(docs []*model.Doc) *pb.Response {
	pbDocs := []*pb.Doc{}
	for _, d := range docs {
		pbDocument := &pb.Doc{
			Id: d.ID,
			Title: d.Title,
			Description: d.Description,
			BulletPoint: d.BulletPoint,
			Brand: d.Brand,
			Color: d.Color,
		}
		pbDocs = append(pbDocs, pbDocument)
	}
	return &pb.Response{Docs: pbDocs}

}

