FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/templates ./templates

ENV GIN_MODE=release
CMD [ "./app" ]
