FROM golang:latest

ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o app ./cmd/main.go
RUN apt-get install dos2unix
CMD ["./app"]
