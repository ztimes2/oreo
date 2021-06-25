FROM golang:1.16.4-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY vendor/ vendor
COPY auth.go .
COPY handlers.go .
COPY http.go .
COPY main.go .

RUN go build -o app -mod vendor *.go

FROM alpine:3.13.5

COPY --from=builder /app/app .

EXPOSE 8081

CMD ./app