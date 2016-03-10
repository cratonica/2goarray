build:
	@go build

test: test-src test-cli

test-cli: build
	@rm -f fixture/sample.go
	@./2goarray < fixture/sample.txt > fixture/sample.go
	@go generate ./fixture
	@go test ./fixture -v

test-src:
	@go test -v -cover

clean:
	@go clean
	@rm -f fixture/sample.go
