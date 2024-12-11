run:
	@go run main.go

build:
	@go build -o desklog-support
	@./desklog-support

watch:
	@air
