# Simple HTTP Server in Go

Este projeto implementa um servidor HTTP simples em Go que lida com três endpoints: `/hello-world`, `/health`, e `/new-endpoint`.

## Endpoints

1. **`/hello-world`**
   - Método: `GET`
   - Descrição: Retorna a mensagem "Hello World!".
   - Exemplo de uso:
     ```bash
     curl -X GET http://localhost:8000/hello-world
     ```

2. **`/health`**
   - Método: `GET`
   - Descrição: Retorna a mensagem "OK!" indicando que o servidor está funcionando corretamente.
   - Exemplo de uso:
     ```bash
     curl -X GET http://localhost:8000/health
     ```

3. **`/new-endpoint`**
   - Método: `GET`
   - Descrição: Retorna a mensagem "OK!".
   - Exemplo de uso:
     ```bash
     curl -X GET http://localhost:8000/new-endpoint
     ```

## Como Executar

1. Certifique-se de ter o Go instalado na sua máquina. Se não tiver, você pode baixar e instalar a partir do site oficial: [https://golang.org/dl/](https://golang.org/dl/).
2. Clone este repositório ou copie o código para um arquivo chamado `main.go`.
3. No terminal, navegue até o diretório onde o arquivo `main.go` está localizado.
4. Execute o comando:
   ```bash
   go run main.go
