build:
	@go build
	@./2goarray --version

test:
	@go test
	@make build

clean:
	go clean
