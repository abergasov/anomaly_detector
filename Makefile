FILE_HASH := $(or ${hash},${hash},"empty_hash")
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

build:
	@echo "-- building gathering binary"
	go build -ldflags "-X main.buildHash=${FILE_HASH} -X main.buildTime=${BUILD_TIME}" -o ./bin/gathering ./cmd/gathering

build_docker:
	@echo "-- building docker binary. buildHash ${FILE_HASH}"
	go build -ldflags "-X main.confFile=common_docker.yml -X main.buildHash=${FILE_HASH} -X main.buildTime=${BUILD_TIME}" -o ./bin/gathering ./cmd/gathering

easy_json:
	@echo "-- generate easy_json"
	@echo "-- remove vendor"
	rm -rf vendor
	@echo "-- generate json"
	~/go/bin/easyjson -all ~/go/src/anomaly_detector/internal/routes/routes_structs.go
	~/go/bin/easyjson -all ~/go/src/anomaly_detector/internal/repository/repository_structs.go
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
