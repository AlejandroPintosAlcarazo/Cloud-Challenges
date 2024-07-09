# Etapa de construccion
FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Construir la aplicación
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/api-go -ldflags "-X main.build=." ./cmd

# Etapa de ejecución
FROM alpine:3.18

COPY --from=builder /app/bin/api-go /app/api-go
WORKDIR /app

RUN chmod +x api-go

EXPOSE 8080

CMD ["./api-go"]
