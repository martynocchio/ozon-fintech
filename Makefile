test:
		go test ./... -cover

build:
		go build -o ./bin/ozon ./cmd/main.go

generate:
		go generate ./...

run_inmemory:
		go run ./cmd/main.go

run_db:
		go run ./cmd/main.go -db
