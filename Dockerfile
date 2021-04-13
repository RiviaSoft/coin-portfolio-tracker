FROM golang:1.16-alpine
WORKDIR /usr/src/
COPY . .
RUN go get ./...
EXPOSE 8000
RUN go build -o ./out ./cmd/server
CMD ["./out"]
