package main

import (
	http "github.com/dbainbriciena/example/internal/pkg/http-integration"
)

func main() {
	integration := http.NewHttpIntegration()
	integration.DoSomething()
}
