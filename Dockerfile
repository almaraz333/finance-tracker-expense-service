FROM golang:1.22.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

EXPOSE 50051

RUN CGO_ENABLED=0 GOOS=linux go build -o /main

CMD ["/main"]
