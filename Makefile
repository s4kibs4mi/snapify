.PHONY: build run_app run_worker build_image

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

build_image:
	docker build -t s4kibs4mi/snapify:latest .
