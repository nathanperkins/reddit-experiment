FROM golang:1 as builder

WORKDIR /app
COPY src/ src/
COPY go.mod .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o top_post_set src/cmd/top_post_set/main.go 

FROM alpine

WORKDIR /app
COPY --from=builder /app/top_post_set .

CMD ["/app/top_post_set"]
