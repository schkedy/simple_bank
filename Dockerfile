# Builds stage
FROM golang:alpine3.22 AS builder
WORKDIR /app
#COPY source dest
COPY . .
RUN go build -o main main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/app.env /app/app.env

EXPOSE 8080
# Expose just documentation which port app listen
CMD ["/app/main"]