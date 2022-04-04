OZON=./bin/ozon

test:
		go test ./... -cover

build:
		go build -o $(OZON) ./cmd/main.go

fmt:
		go fmt ./...
		goimports -l ./
		go mod tidy

generate:
		go generate ./...

run_im:
		$(OZON)

run_db:
		$(OZON) -db
