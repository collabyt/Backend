FROM golang:latest

LABEL maintainer="Gustavo H. M. Silva <gushmsilva@pm.me>"

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

ENV PORT 8080
ENV DBHOST 0.0.0.0
ENV DBUSER postgres
ENV DBPASSWORD postgres
ENV DBDATABASE collabyt
ENV DBPORT 5432
ENV DBSSL disable
ENV DBSOURCE postgres

RUN go build
CMD ["./Backend"]