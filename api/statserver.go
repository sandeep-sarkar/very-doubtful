package api

import (
	"context"
	"log"

	api "github.com/very-doubtful/proto/calcstatisticsb"
)

type Server struct {
	api.UnimplementedStatisticsCalculatorServer
}

func (s *Server) CalculateStatistics(
	ctx context.Context,
	req *api.CalculateStatisticsRequest,
) (*api.CalculateStatisticsResponse, error) {
	log.Printf("Received document...")
	log.Printf("Processing document")

	sc := StatCalculator{
		Document:       req.Document.GetContent(),
		ColumnsExclude: req.GetColumnsExclude(),
		ColumnsInclude: req.GetColumnsInclude(),
		PrimaryColumn:  req.GetPrimaryColumn(),
	}

	content, err := sc.calculateStatistics()
	if err != nil {
		content = []byte("Error in processing document")
	}

	/**
	return &api.CalculateStatisticsResponse{
		Content: req.Document.GetContent(),
	}, nil
	**/

	return &api.CalculateStatisticsResponse{
		Content: content,
	}, nil
}
