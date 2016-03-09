build:
	@go build

test: build
	@rm -f fixture/sample.go
	@./2goarray < fixture/sample.txt > fixture/sample.go
	@go generate ./fixture
	@go test ./fixture

clean:
	@go clean
	@rm -f fixture/sample.go
