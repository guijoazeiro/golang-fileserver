# Basic Auth File Server

Este é um simples servidor de arquivos HTTP com autenticação básica.

## Requisitos

- Go 1.16 ou superior
- [go-http-auth](https://github.com/abbot/go-http-auth)

## Instalação

1. Clone o repositório:

```bash
git clone https://github.com/guijoazeiro/golang-fileserver.git
cd golang-fileserver
```
## Instale as dependências
```bash
go get github.com/abbot/go-http-auth
```

## Uso
Compile e execute o servidor com o seguinte comando, fornecendo o diretório HTTP e a porta como argumentos:

```bash
go run main.go <http-dir> <port>
```

Por exemplo:

```bash
go run cmd/fileserver/main.go static 8080
```
Isso iniciará o servidor na porta 8080 e servirá arquivos do diretório /static
