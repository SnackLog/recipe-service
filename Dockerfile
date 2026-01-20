FROM golang:alpine AS build

WORKDIR /app

ENV GONOPROXY=github.com/SnackLog/*
ENV GOSUMDB=off

RUN apk add git

COPY go.mod .
RUN go mod download

COPY . .
RUN go build -o /app/recipe-service

FROM alpine:latest
ENV GIN_MODE=release

RUN apk add curl

WORKDIR /app
COPY LICENSES /licenses
COPY --from=build /app/recipe-service .

ENTRYPOINT ["./recipe-service"]

