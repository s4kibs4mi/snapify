.PHONY: build run_app run_worker

export GO111MODULE=on
export CGO_ENABLED=0
export CONFIG_FILE=./config.yml

build:
	go build -o ./bin/app ./cmd/app
	go build -o ./bin/worker ./cmd/worker

run_app:
	./bin/app

run_worker:
	./bin/worker
