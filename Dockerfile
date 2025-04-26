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

RUN apk --no-cache add ca-certificates tzdata

RUN adduser -D -u 10001 appuser
USER appuser

WORKDIR /app

COPY --from=build /app/server .

EXPOSE 3000

ENTRYPOINT ["./server"]