# Usa la imagen oficial de Go como base
FROM golang:1.20

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el código fuente de la aplicación
COPY . .

# Inicializa un módulo Go en el directorio actual
RUN go mod init my-go-app

RUN go mod tidy
# Compila la aplicación
RUN go build -o main .

# Define el comando predeterminado para ejecutar el binario compilado
CMD ["./main"]

