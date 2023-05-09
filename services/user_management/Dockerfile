FROM golang:1.20-alpine

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY go.mod go.sum ./
COPY ./services/user_management /app/services/user_management
COPY ./pkg /app/pkg
COPY ./app.env /app

RUN go mod tidy
RUN go mod download

RUN go build -o app /app/services/user_management/cmd/main.go

EXPOSE 3007
