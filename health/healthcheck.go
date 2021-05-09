package health

import (
	"golang.org/x/net/context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type HealthCheckService struct{}

// Check is a part of grpc health check service
// https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto
func (s *HealthCheckService) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (
	*grpc_health_v1.HealthCheckResponse, error) {

	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

// Watch is a part of grpc health check service
// https://github.com/grpc/grpc/blob/master/src/proto/grpc/health/v1/health.proto
// NOTE: Watch is currently not implemented!
func (s *HealthCheckService) Watch(in *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {

	return status.Error(codes.Unimplemented, "Watching is not supported")
}

func NewHealthCheckService() *HealthCheckService {
	return &HealthCheckService{}
}
