package main

import (
	grpc "github.com/dbainbriciena/example/internal/pkg/grpc-integration"
)

func main() {
	integration := grpc.NewGrpcIntegration()
	integration.DoSomething()
}
