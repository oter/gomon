FROM golang:1.15.2 AS build

RUN apt-get update && apt-get install -y --no-install-recommends

WORKDIR /src
COPY . .

COPY go.mod go.sum tools.go Makefile main.go ./
COPY ./cmd ./cmd
COPY ./pkg ./pkg

RUN go mod download

RUN make all
