build:
	@go build -o bin/goai

run: build
	@./bin/goai

test:
	go test -v ./... -count=1