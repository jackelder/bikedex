test:
	go test ./... -cover

run:
	. ./scripts/set-environment.sh && \
	go run main.go