build:
	@go build
	@./2goarray --version

test:
	@go test -v
	@make build

clean:
	go clean
