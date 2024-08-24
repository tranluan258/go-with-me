FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go 

EXPOSE 8080

CMD ["./main"]
