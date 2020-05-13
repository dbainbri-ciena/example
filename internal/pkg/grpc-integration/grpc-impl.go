package integration

import (
	"fmt"

	"github.com/dbainbriciena/example/internal/pkg/common"
)

func init() {
	fmt.Println("GRPC INTEGRATION LOADED")
}

type grpcImpl struct{}

func NewGrpcIntegration() common.Integration {
	return &grpcImpl{}
}

func (k *grpcImpl) DoSomething() error {
	fmt.Println("GRPC")
	return nil
}
