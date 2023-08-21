FROM golang:1.21.0-alpine3.18 AS builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY .env .
RUN go mod download
COPY . .
RUN go build -o main main.go

FROM alpine:3.18
WORKDIR /app
COPY .env .
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]