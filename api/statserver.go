package api

import (
	"context"
	"fmt"

	api "github.com/very-doubtful/proto/calcstatisticsb"
)

type Server struct {
	api.UnimplementedStatisticsCalculatorServer
}

func (s *Server) CalculateStatistics(
	ctx context.Context,
	req *api.CalculateStatisticsRequest,
) (*api.CalculateStatisticsResponse, error) {
	fmt.Println("Received document...")
	fmt.Println(req.Document.GetContent())

	return &api.CalculateStatisticsResponse{
		Content: req.Document.GetContent(),
	}, nil
}
