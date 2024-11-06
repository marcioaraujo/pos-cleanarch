# Desafio 3 da Pós-graduação Full Cycle Go Expert

## Enunciado do desafio

Para este desafio, você precisará criar o usecase de listagem das orders.
Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)
- Service ListOrders com GRPC
- Query ListOrders GraphQL
  Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.

Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.

Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

## Instruções de desenvolvimento

Para regerar fontes da injeção de dependências (Wire):
https://github.com/google/wire

```bash
./make_wire.sh
```

Para regerar fontes GraphQL:
https://github.com/graphql-go/graphql

```bash
./make_gql.sh
```

Para regerar fontes gRPC:
https://grpc.io/docs/languages/go/quickstart/

```bash
./make_grpc_proto.sh
```

Para rodar migrations:

```
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

## Instruções de uso

Inicializar dependências da aplicação (Banco de dados MySQL e Broker de mensageria RabbitMQ)

```bash
docker compose up -d
```

Executar aplicação

```bash
./init.sh
```

Portas padrão:

- Servidor HTTP: 8000
- Servidor gRPC: 50051
- Servidor GraphQL: 8080
