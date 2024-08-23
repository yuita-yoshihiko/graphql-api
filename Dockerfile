FROM golang:1.23.0-alpine

WORKDIR /graphql-api
COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git alpine-sdk
RUN set -x \
    && go mod download \
    && go install github.com/rubenv/sql-migrate/...@latest \
    && go install github.com/volatiletech/sqlboiler/v4@latest \
    && go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest \
    && go install github.com/air-verse/air@latest \
    && go install github.com/99designs/gqlgen@latest

COPY . .
