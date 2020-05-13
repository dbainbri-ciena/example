package main

import (
	kafka "github.com/dbainbriciena/example/internal/pkg/kafka-integration"
)

func main() {
	integration := kafka.NewKafkaIntegration()
	integration.DoSomething()
}
