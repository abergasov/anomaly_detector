FILE_HASH := $(or ${hash},${hash},"empty_hash")
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	@echo "-- building gathering binary"
	go build -ldflags "-X main.buildHash=${FILE_HASH} -X main.buildTime=${BUILD_TIME}" -o ./bin/gathering ./cmd/gathering
	go build -ldflags "-X main.buildHash=${FILE_HASH} -X main.buildTime=${BUILD_TIME}" -o ./bin/analyser ./cmd/analyser

build_docker:
	@echo "-- building docker binary. buildHash ${FILE_HASH}"
	go build -ldflags "-X main.buildHash=${FILE_HASH} -X main.buildTime=${BUILD_TIME}" -o ./bin/gathering ./cmd/gathering
	go build -ldflags "-X main.buildHash=${FILE_HASH} -X main.buildTime=${BUILD_TIME}" -o ./bin/analyser ./cmd/analyser

easy_json:
	@echo "-- generate easy_json"
	@echo "-- remove vendor"
	rm -rf vendor
	@echo "-- generate json"
	easyjson -all ~/go/src/anomaly_detector/internal/routes/gathering/gathering_structs.go
	easyjson -all ~/go/src/anomaly_detector/internal/routes/analyser/analyser_structs.go
	easyjson -all ~/go/src/anomaly_detector/internal/repository/repository_structs.go
	@echo "-- restore vendor"
	go mod vendor

gogen:
	@echo "-- generate code"
	go generate ./internal...

lint:
	@echo "-- format code"
	gofmt -s -w .
	@echo "-- check code"
	golangci-lint run -c .golangci.yaml ./internal...
	golangci-lint run -c .golangci.yaml ./cmd...

test:
	@echo "-- test code"
	go test ./internal...