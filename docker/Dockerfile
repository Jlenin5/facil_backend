FROM golang:1.23.4-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

# Copiar módulos y descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código fuente
COPY . .

# Exponer el puerto de la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["air", "-c", ".air.toml"]