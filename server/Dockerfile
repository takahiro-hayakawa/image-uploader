FROM golang:1.14

ENV GO111MODULE=on
WORKDIR /go/src

COPY go.mod go.sum ./
COPY main.go ./

RUN go mod download

RUN go build -o app main.go
CMD ["./app"]
