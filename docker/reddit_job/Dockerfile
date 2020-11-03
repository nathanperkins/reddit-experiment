FROM golang:1 as builder

WORKDIR /app
COPY src/ src/
COPY go.mod .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o reddit src/cmd/reddit_job/main.go 

FROM alpine

WORKDIR /app
COPY --from=builder /app/reddit .

CMD ["/app/reddit"]
