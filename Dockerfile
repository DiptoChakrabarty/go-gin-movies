FROM golang:1.15-alpine AS builder
RUN apk update && apk add --update gcc g++
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o app main.go
RUN go get -u github.com/swaggo/swag/cmd/swag && /go/bin/swag init

FROM alpine:latest
RUN apk --no-cache add ca-certificates mailcap && addgroup -S app && adduser -S app -G app -u 101
USER app
WORKDIR /app
RUN mkdir db
COPY --chown=app:app --from=builder /app/app ./
COPY --chown=app:app --from=builder /app/docs ./docs
EXPOSE 8000
ENV GIN_MODE=release
CMD ["./app"]
