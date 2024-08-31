FROM golang:1.23.0-alpine3.20 AS build

ENV GO111MODULE=on

WORKDIR /bin

COPY go.mod go.sum /bin/
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /billing

EXPOSE 9000

CMD ["/billing"]