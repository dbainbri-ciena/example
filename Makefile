build: kafka grpc http

check: build
	@echo "Validate Kafka binary does not contain HTTP or GRPC packages"
	@test $(shell strings ./bin/kafka-cmd | grep -E -c '(http|grpc)') -eq 0
	@echo "Validate GRPC binary does not contain HTTP or Kafka packages"
	@test $(shell strings ./bin/grpc-cmd | grep -E -c '(http|kafka)') -eq 0
	@echo "Validate HTTP binary does not contain Kafka or GRPC packages"
	@test $(shell strings ./bin/http-cmd | grep -E -c '(kafka|grpc)') -eq 0

run: build
	@echo "RUN Kafka"
	./bin/kafka-cmd
	@echo "=====\n\n"
	@echo "RUN GRPC"
	./bin/grpc-cmd
	@echo "=====\n\n"
	@echo "RUN HTTP"
	./bin/http-cmd
	@echo "=====\n\n"

kafka:
	go build -o bin/kafka-cmd ./cmd/kafka-cmd/...

grpc:
	go build -o bin/grpc-cmd ./cmd/grpc-cmd/...

http:
	go build -o bin/http-cmd ./cmd/http-cmd/...

clean:
	rm -rf bin
