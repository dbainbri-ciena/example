package integration

import (
	"fmt"

	"github.com/dbainbriciena/example/internal/pkg/common"
)

func init() {
	fmt.Println("HTTP INTEGRATION LOADED")
}

type httpImpl struct{}

func NewHttpIntegration() common.Integration {
	return &httpImpl{}
}

func (k *httpImpl) DoSomething() error {
	fmt.Println("HTTP")
	return nil
}
