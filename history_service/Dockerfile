ARG GO_VERSION=1.21
FROM golang:${GO_VERSION} AS build
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . ./
RUN go build -v -o history-service ./cmd/
CMD ["/app/history-service"]