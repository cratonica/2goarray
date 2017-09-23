build:
	@go build

test: build
	@./2gobytes --version

clean:
	@go clean
