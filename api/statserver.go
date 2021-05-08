package api

import (
	"context"
	"fmt"

	api "github.com/very-doubtful/proto"
)

type StatServer struct {
	api.UnimplementedCalculateStatisticsServer
}

func (s *StatServer) CalculateStatistics(
	ctx context.Context,
	req *api.CalculateStatisticsRequest,
) (*api.CalculateStatisticsResponse, error) {
	fmt.Println("Received document...")
	fmt.Println(req.Document.GetContent())

	return &api.CalculateStatisticsResponse{
		Content: req.Document.GetContent(),
	}, nil
}
