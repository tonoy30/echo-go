# Start from the latest golang base image
FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

## Add this go mod download command to pull in any dependencies
RUN go mod download

## Our project will now successfully build with the necessary go libraries included.
RUN go build -o ./bin/main ./cmd/api/main.go

EXPOSE 5050

## our newly created binary executable
CMD ["/app/bin/main"]

