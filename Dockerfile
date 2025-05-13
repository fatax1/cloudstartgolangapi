# Steg 1: Bygg Go-bilden
FROM golang:alpine AS builder

# Installera nödvändiga paket
RUN apk update && apk add --no-cache git

# Ställ in arbetskatalogen till /app
WORKDIR /app

# Kopiera go.mod och go.sum och kör go mod tidy för att hämta alla beroenden
COPY go.mod go.sum ./
RUN go mod tidy

# Kopiera resten av källkoden
COPY . .

# Kör testkommandon
RUN go test ./...

# Bygg applikationen
RUN go build -o /app/cmd/site .

# Steg 2: Skapa den slutliga Docker-bilden
FROM alpine:latest

# Skapa arbetsmapp och kopiera den byggda binären från builder-fasen
WORKDIR /root/
COPY --from=builder /app/cmd/site /site

# Om du behöver yml-filer
COPY *.yml / 

# Exponera port 8080
EXPOSE 8080

# Kör applikationen
ENTRYPOINT ["/site"]
