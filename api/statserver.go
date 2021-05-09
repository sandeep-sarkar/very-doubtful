package api

import (
	"context"
	"errors"
	"log"

	api "github.com/very-doubtful/proto/calcstatisticsb"
	"google.golang.org/grpc/peer"
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
	}

	documentSource, err := sc.calculateStatistics()
	if err != nil {
		return nil, errors.New("Error in processing document")
	}

	peer, _ := peer.FromContext(ctx)

	log.Printf("Served request to %s", peer.Addr.String())

	return &api.CalculateStatisticsResponse{
		DocumentSource: &api.DocumentSource{
			HttpUri: documentSource,
		},
	}, nil
}
