FROM golang:1.24-alpine

WORKDIR /github-api
COPY . .
RUN mkdir -p ./bin && go build -o ./bin/main ./cmd/api
CMD ["./bin/main"]
