FROM golang:1.22-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o product-service

EXPOSE 8080

CMD ["./product-service"]
