FROM golang:alpine AS builder

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
COPY expense.db ./
RUN CGO_ENABLED=1 GOOS=linux go build -tags netgo -ldflags '-w -extldflags "-static"' -o /main ./

FROM scratch

WORKDIR /root

COPY --from=builder /main .
COPY --from=builder /app/expense.db /root/

EXPOSE 50051

CMD ["/root/main"]
