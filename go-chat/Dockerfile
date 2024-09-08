FROM golang:1.22-alpine AS deps

WORKDIR /app

COPY go.mod go.sum ./
COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build cmd/main.go 

FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY --from=deps /app/main ./

COPY --from=deps /app/.env ./
COPY --from=deps /app/firebase-admin.json ./
COPY --from=deps /app/views ./views

EXPOSE 8080

CMD ["./main"]
