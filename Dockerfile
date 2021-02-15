FROM golang:latest

LABEL maintainer="Gustavo H. M. Silva <gushmsilva@pm.me>"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV CORS_ADDRESS localhost
ENV CORS_PORT 9000
ENV APP_ADDRESS 0.0.0.0
ENV APP_PORT 8080
ENV APP_IDLE_TIMEOUT 120
ENV APP_READ_TIMEOUT 5
ENV APP_WRITE_TIMEOUT 5

ENV CACHE_TTL 60
ENV CACHE_HOST localhost
ENV CACHE_PORT 6379
ENV CACHE_PASSWORD ""

RUN go build ./cmd/collabyt/main.go
CMD ["./main"]

EXPOSE 8080