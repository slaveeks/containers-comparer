# Build stage
FROM golang:1.16 AS build

WORKDIR /app

COPY . .

RUN go build -o main .

# Final stage
FROM debian:buster-slim

WORKDIR /app

COPY --from=build /app/main .

CMD ["./main"]
