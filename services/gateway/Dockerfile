FROM golang:1.20-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum ./
COPY ./services/gateway /app/services/gateway
COPY ./pkg /app/pkg
COPY ./app.env /app

RUN go mod tidy
RUN go mod download

RUN go build -o app /app/services/gateway/cmd/main.go

EXPOSE 8080
