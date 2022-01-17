run:
	go run main.go

dev:
	CGO_ENABLED=0 go build -ldflags="-w -s" -o  build/main main.go
	./build/main



critic:
	gocritic check -enableAll ./...
