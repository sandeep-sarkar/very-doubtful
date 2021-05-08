package main

import (
	"context"
	"io/ioutil"
	"log"

	pb "github.com/very-doubtful/proto/calcstatisticsb"

	"google.golang.org/grpc"
)

const (
	host = "localhost:50061"
)

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewStatisticsCalculatorClient(conn)

	document, err := ioutil.ReadFile("test.csv")
	if err != nil {
		log.Fatal("Couldn't read input document")
	}
	ctx := context.Background()
	resp, err := client.CalculateStatistics(ctx, &pb.CalculateStatisticsRequest{
		Document: &pb.Document{
			Content: document,
		},
		ColumnsExclude: []string{"aaa", "bbb"},
	})

	ioutil.WriteFile("out.csv", resp.GetContent(), 0644)
}
