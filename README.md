# GoTransfer

Serviço para realizar tranferência com liquidação.

## 🚀 Começando

Essas instruções permitirão que você obtenha uma cópia do projeto em operação na sua máquina local para fins de desenvolvimento e teste.

Consulte **Pré-requisitos** e **Instalação** para saber como implantar o projeto.

### 📋 Pré-requisitos

De que coisas você precisa para instalar o software e como instalá-lo?

```
Para Rodar em Container:
    DOCKER (https://www.docker.com/)

Para Rodar em Local:
    GO (https://go.dev/)
    MongoDB (https://www.mongodb.com/)
    RabbitMQ (https://www.rabbitmq.com)

```

### 🔧 Instalação

São duas as formas de execução, dentro de um container (Docker) ou local.

DOCKER:

```
Uma vez com o docker instalado, abrir terminal na raiz do projeto e executar:

    docker-compose up --build --force-recreate    
```

LOCAL:

```
Uma vez os **Pré-requisitos** instalados, executar na raiz do projeto:
    * obs: substituir os valoresm em [] pelo valor correto
    PRESENTER=api MONGO_URL=[MONGO_URL] RABBIT_URL=[RABBIT_URL] LIQUIDATION_API_URL=[LIQUIDATION_API_URL] go run main.go
```

A api irá responder na porta 8000 e com documentação swagger no path /swagger/index.html, Ex: http://localhost:8000/swagger/index.html

## ⚙️ Executando os testes

Para execução dos testes basta executar o comando:
    go test ~/github.com/gabrielnando1/go_service_transfer/tests/use_cases/bo



## 🛠️ Construído com

* [Go](https://go.dev/) - Uma das linguagens que mais vem creescedo
* [Gin](https://github.com/gin-gonic/gin) - O framework go mais utilizado para criação de apis
* [MongoDB](https://www.mongodb.com/) - Banco de dados NoSQL 
* [RabbitMQ](https://www.rabbitmq.com) - Servidor de mensageria open-source


## ✒️ Autores


* **Gabriel Ribeiro** - *Trabalho Inicial* - [desenvolvedor](https://github.com/gabrielnando1)

