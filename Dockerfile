FROM golang:latest

LABEL maintainer="Gustavo H. M. Silva <gushmsilva@pm.me>"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV APP_PORT 8080
ENV APP_IDLE_TIMEOUT 120
ENV APP_READ_TIMEOUT 5
ENV APP_WRITE_TIMEOUT 5

ENV DB_HOST localhost
ENV DB_USER postgres
ENV DB_PASSWORD postgres
ENV DB_DATABASE collabyt
ENV DB_PORT 5432
ENV DB_SSL disable
ENV DB_SOURCE postgres

ENV CACHE_TTL 60
ENV CACHE_HOST localhost
ENV CACHE_PORT 6379
ENV CACHE_PASSWORD ""

RUN go build
CMD ["./Backend"]