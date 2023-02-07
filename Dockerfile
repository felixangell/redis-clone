FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build cmd/server/main.go -o /main

EXPOSE 9093

CMD [ "/main" ]