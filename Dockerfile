FROM golang:alpine as builder
WORKDIR /application
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o application ./cmd/main.go

FROM alpine:latest
WORKDIR /usr/bin
COPY --from=builder /application/application .
CMD ["application"]