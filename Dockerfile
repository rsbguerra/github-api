FROM golang:1.24-alpine

WORKDIR /github-api
COPY . .
RUN go build -o main .
CMD ["./main"]
