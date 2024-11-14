# Clean Architecture - Order Management System

## Desafio Proposto

Bem-vindo, Devs!

Chegou a hora de colocar em prática seus conhecimentos. O desafio consiste em criar o caso de uso para a listagem das ordens (`Orders`). Essa listagem deve ser implementada utilizando:

- **Endpoint REST:** `GET /order`
- **Serviço gRPC:** `ListOrders`
- **Query GraphQL:** `ListOrders`

Além disso, não se esqueça de criar as migrações necessárias e o arquivo `api.http` contendo as requisições para criar e listar as ordens.

Para a criação do banco de dados, utilize o Docker. O comando `docker compose up` deverá configurar e iniciar tudo automaticamente, incluindo o banco de dados.

Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

# Como rodar o projeto:

1. **Clone o repositório:**

   ```bash
   git clone https://github.com/marcioaraujo/pos-cleanarch.git && cd pos-cleanarch
   ```

2. **Instruções para boot:**

   ```bash
   docker compose up -d
   ```

   > **Nota:** Na plataforma Apple M3 apresentou problemas de compatibilidade de containers arm64/amd64, resultando indisponibilidade do rabbitmq ou mysql.

   > **Nota:** Tudo se resolveu com reinicio do docker:

   ```bash
   docker-compose down && docker-compose up -d
   ```

# Validação do Desafio

Os serviços são acessados nestas portas:

- **Web:** `localhost:8000`
- **gRPC:** `localhost:50051`
- **GraphQL:** `localhost:8080`

### Teste

#### WebServer

Na pasta api estão os arquivos para execução das chamadas api Rest `api/create_order.http` e `api/list_order.http`.

> **Nota:** Caso queira, instale a extensão REST Client no vscode, para facilitar a exeção de arquivos .http

#### gRPC

Para testar via gRPC, utilize o [Evans](https://github.com/ktr0731/evans):

> **Nota:** Na plataforma Apple pode ser instalado pelo seguinte comando:

```bash
    brew install evans
```

1. Consumindo o serviço:

   ```bash
   evans -r repl --host 127.0.0.1 --port 50051
   ```

2. Vá até package:

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
     createOrder(input: { id: "unique_id", Price: 18.50, Tax: 3.2 }) {
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
