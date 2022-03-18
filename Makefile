dev:
	nodemon --exec go run $(shell pwd)/cmd/main.go --signal SIGTERM