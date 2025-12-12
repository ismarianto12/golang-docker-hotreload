FROM golang:1.24-alpine

RUN apk update && apk add --no-cache git bash

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Install CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon@latest

# Copy source code
COPY . .

EXPOSE 6060

# Gunakan CompileDaemon untuk development (hot reload)
CMD ["CompileDaemon", "-directory=/app", "-build=go build -o /tmp/app cmd/main.go", "-command=/tmp/app"]