.phony: build clean deploy

clean:
	@go clean
	@rm -rf ./bin

build:
	./gobuild.sh

# TODO: for some reason, the user docker flag doesn't seem to work
start:  build
	sls offline --useDocker start  --host 0.0.0.0

format:
	gofmt -s -w .