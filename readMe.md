# Go Expert Challenge - Rate Limiter

Implementação de observabilidade utilizando Opentelemetry e zipkin para monitorar a execução dos serviços que buscam a informação da localização e temperatura do usuario a partir do cep enviado pelo query string

## Arquitetura

A aplicação é composta se 2 serviços, o primeiro, o serviço cep, é composto por um servidor web e um client gRPC. Já o segundo, o serviço weather, é composto apenas pelo servidor gRPC.
![alt text](img/application.png)

## Fluxograma

O servidor web que recebe requisições HTTP, valida se o CEP enviado pelo usuário é válido e caso seja, realiza uma chamada http externa para o [via cep](https://viacep.com.br/) para encontrar os dados da cidade e com o nome da cidade, encontra a temperatura atual realizando uma chamada externa para o [wather api](https://www.weatherapi.com/), sendo o monitoramento dese fluxo realizado utilizando [Opentelemetry](https://opentelemetry.io/) e [Zipkin](https://zipkin.io/) para o trace.

## Executando o projeto

**Obs:** é necessário ter o [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/) instalados.

1. Crie um arquivo `.env` na raiz de cada projeto copiando o conteúdo de `.env.example` e ajuste-o conforme necessário. Por padrão, os seguintes valores são utilizados:

```sh
# cep
GRPC_PORT=50051
GRPC_SERVER_NAME=weather
HTTP_PORT=8080
REQUEST_NAME_OTEL=microservice-cep-request
OTEL_SERVICE_NAME=microservice-cep
OTEL_EXPORTER_OTLP_ENDPOINT=otel-collector:4317

# weather
GRPC_PORT=50051
WEATHER_API_KEY=XXXX
REQUEST_NAME_OTEL=microservice-weather-request
OTEL_SERVICE_NAME=microservice-weather
OTEL_EXPORTER_OTLP_ENDPOINT=otel-collector:4317
```

2. Rode o comando

```
docker-compose up
```

### Request

| Endpoint | Descrição                                 | Método | Parâmetro |
| -------- | ----------------------------------------- | ------ | --------- |
| /        | Calcula a temperatura atual em uma cidade | GET    | zipcode   |

**Requisição:**

```sh
$ curl -X GET http://localhost:8080/?zipcode=01153000

# Response
{"temp_C":28.100000381469727,"temp_F":82.5999984741211,"temp_K":301.1000061035156}
```
**Zepkin:**
http://localhost:9411/zipkin/traces
![alt text](img/zepkin.png)