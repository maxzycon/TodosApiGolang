FROM golang:1.18.1-alpine3.15 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.15
WORKDIR /app
COPY wait-for.sh .
COPY --from=builder /app/main .

EXPOSE 8080
CMD [ "/app/main" ]
