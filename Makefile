build:
	@go build

test: test-src test-cli

test-src:
	@go test -v -cover ./generator

test-cli: clean build
	@rm -f fixture/sample.go
	@./2gobytes < fixture/sample.txt > fixture/sample.go
	@go generate ./fixture
	@go test ./fixture -v
	@go generate ./fixture/sample_files
	@go run ./fixture/sample_files/*.go


clean:
	@go clean
	@rm -f fixture/sample.go
