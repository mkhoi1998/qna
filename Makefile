build:
	go build -o api *.go

run: build
	./api; rm api
