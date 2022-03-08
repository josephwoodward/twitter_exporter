build:
	env GOOS=linux GOARCH=amd64 go build -o twitter_exporter cmd/main.go