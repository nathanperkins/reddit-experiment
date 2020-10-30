FROM golang:1 as builder

WORKDIR /go/src/app
COPY src/ src/
COPY go.mod .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 go build -o reddit /go/src/app/src/cmd/reddit_job/main.go 

FROM alpine

WORKDIR /app
COPY --from=builder /go/src/app/reddit .

CMD ["/app/reddit"]
