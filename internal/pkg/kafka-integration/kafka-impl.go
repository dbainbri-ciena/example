package integration

import (
	"fmt"

	"github.com/dbainbriciena/example/internal/pkg/common"
)

func init() {
	fmt.Println("KAFKA INTEGRATION LOADED")
}

type kafkaImpl struct{}

func NewKafkaIntegration() common.Integration {
	return &kafkaImpl{}
}

func (k *kafkaImpl) DoSomething() error {
	fmt.Println("KAFKA")
	return nil
}
