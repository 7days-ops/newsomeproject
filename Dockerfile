FROM golang:1.24.2 AS builder

WORKDIR /app

COPY . .

RUN go build -o myapp .

FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/myapp .

CMD [ "./myapp" ]

