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

## Instruções de desenvolvimento e teste

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

```bash
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

ou

```bash
./make_migrate.sh

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

ou

```bash
cd cmd/ordersystem
go run main.go wire_gen.go
```

Portas padrão:

- Servidor HTTP: 8000
- Servidor gRPC: 50051
- Servidor GraphQL: 8080

### Testando a Aplicação

#### WebServer

As requisições REST para criar e listar ordens podem ser testadas utilizando os arquivos `api/create_order.http` e `api/list_order.http`.

#### gRPC

Para testar via gRPC, utilize o [Evans](https://github.com/ktr0731/evans):

1. Conecte-se ao serviço:

   ```bash
   evans -r repl --host 127.0.0.1 --port 50051
   ```

2. Navegue até o package:

   ```bash
   package pb
   ```

3. Acesse o serviço:

   ```bash
   service OrderService
   ```

4. Chame as funções:

   - **Criar Ordem:** `call CreateOrder`
   - **Listar Ordens:** `call ListOrders`

#### GraphQL

Para acessar o serviço GraphQL:

1. Abra o navegador e acesse: `http://localhost:8080/`
2. Digite a mutation e a query para rodar as operações `createOrder` e `queryOrder`:

   ```graphql
   mutation createOrder {
     createOrder(input: { id: "xxxxx", Price: 10.2, Tax: 2.0 }) {
       id
       Price
       Tax
       FinalPrice
     }
   }

   query queryOrder {
     order {
       id
       Price
       Tax
       FinalPrice
     }
   }
   ```
