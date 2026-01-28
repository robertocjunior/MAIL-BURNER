# Estágio de Compilação
FROM golang:1.24-alpine AS builder

# Instalar dependências necessárias para o SQLite (CGO)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Copiar arquivos de dependências primeiro para cachear camadas
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante do código
COPY . .

# Compilar o binário (CGO_ENABLED=1 é necessário para o driver SQLite)
RUN CGO_ENABLED=1 GOOS=linux go build -o tempmail ./cmd/tempmail/main.go

# Estágio Final (Execução)
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /root/

# Copiar o binário do estágio anterior
COPY --from=builder /app/tempmail .

# Copiar a pasta static (necessária para o servidor web)
COPY --from=builder /app/static ./static

# Expor a porta padrão definida no seu config.go
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./tempmail"]