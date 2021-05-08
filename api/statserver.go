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

	sc := StatCalculator{
		Document:       req.Document.GetContent(),
		ColumnsExclude: req.GetColumnsExclude(),
		ColumnsInclude: req.GetColumnsInclude(),
		PrimaryColumn:  req.GetPrimaryColumn(),
	}

	sc.printDocument()
	fmt.Println("Columns excluded", sc.ColumnsExclude)

	return &api.CalculateStatisticsResponse{
		Content: req.Document.GetContent(),
	}, nil
}
