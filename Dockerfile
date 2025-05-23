FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

# Default CMD just runs with http mode; can be overridden
CMD ["./main", "--svc=http"]
