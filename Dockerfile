# syntax=docker/dockerfile:1

ARG GO_VERSION=1.24.2
FROM golang:${GO_VERSION} AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Final lightweight image
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata postgresql-client

WORKDIR /app

COPY --from=build /app/server .
COPY wait-for-db.sh .

# Important: chmod before switching user
RUN chmod +x wait-for-db.sh

RUN adduser -D -u 10001 appuser
USER appuser

EXPOSE 3000

ENTRYPOINT ["./wait-for-db.sh", "./server"]