.phony: build clean deploy

clean:
	@go clean
	@rm -rf ./bin

build:
	./gobuild.sh

start: build
	sls offline --useDocker start  --host 0.0.0.0

format:
	gofmt -s -w .