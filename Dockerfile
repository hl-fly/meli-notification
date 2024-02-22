# Use uma imagem base que tenha o Go instalado
FROM golang:latest

# Defina o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copie os arquivos do código-fonte para o diretório de trabalho
COPY . .

# Baixe as dependências do módulo Go
RUN go mod download

# Compile o aplicativo Go
RUN go build -o app

# Exponha a porta em que o aplicativo irá escutar
EXPOSE 8080

# Defina o comando de inicialização do aplicativo
CMD ["./app"]
