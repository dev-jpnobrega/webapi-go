# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:1.13-alpine

LABEL maintainer="JP <joaopaulonobrega@nobrega.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./src

EXPOSE 8080

CMD ["./main"]