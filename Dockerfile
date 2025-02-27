# Usar la imagen de golang con la versión 1.23
FROM golang:1.23-alpine

# Establecer el directorio de trabajo en el contenedor
WORKDIR /app

# Copiar los archivos de go.mod y go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo el código de la aplicación
COPY . .

# Crear el directorio de la aplicación para almacenar los archivos de la BD
RUN mkdir -p /app && chmod 777 /app

# Compilar la aplicación
RUN go build -o main .

# Ejecutar la aplicación
CMD ["./main"]

# Exponer el puerto 8080
EXPOSE 8080