FROM golang:1.20.5

WORKDIR /app
COPY . .
RUN go mod download

COPY ./pkg/config/config.sample.yaml ./pkg/config/config.yaml
RUN go build -o /app/simpleshop ./cmd/api

EXPOSE 8888
ENTRYPOINT ["/app/simpleshop"]
