# Etapa de construcción
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -tags connector -o ./bin/conector-aemet -ldflags "-X main.build=." ./cmd/updater/main.go


# Etapa de ejecución
FROM alpine:3.18


COPY --from=builder /app/bin/conector-aemet /app/conector-aemet
COPY gcloud-key.json /app/gcloud-key.json

WORKDIR /app

RUN chmod +x conector-aemet

ENV GOOGLE_APPLICATION_CREDENTIALS="/app/gcloud-key.json"

EXPOSE 8080

CMD ["./conector-aemet"]