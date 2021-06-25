FROM golang:1.16.4-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY vendor/ vendor
COPY main.go .

RUN go build -o app -mod vendor *.go

FROM alpine:3.13.5

COPY --from=builder /app/app .
COPY static/ static

EXPOSE 8082

CMD ./app