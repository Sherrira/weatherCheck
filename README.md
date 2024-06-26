# weatherCheck
Desafio - Sistema de temperatura por CEP - Go Express

## Índice
1. [Pré-requisitos](#pré-requisitos)
2. [Configurações de ambiente](#configurações-de-ambiente)
3. [Executando o projeto](#executando-o-projeto)
4. [Interagindo com a API](#interagindo-com-a-api)
5. [Consultando temperatura por CEP](#consultando-temperatura-por-cep)

## Pré-requisitos
Assegure-se de ter as seguintes ferramentas instaladas:
- [Golang](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/compose/install/)

## Configurações de ambiente
Para executar as consultas de temperatura é necessária a obtençao de uma chave de acesso à API WeatherAPI. Para isso, siga os passos abaixo:    
1. Acesse o site [WeatherAPI](https://www.weatherapi.com/) e crie uma conta.
2. Após a criação da conta, acesse o painel de controle e copie a chave de acesso.
3. No arquivo ".env" na pasta cmd, adicione a chave de acesso obtida no passo anterior.

Nota: o projeto está entregando um arquivo Dockerfile que permite o deploy na Google Cloud Platform no serviço Google Cloud Run.

## Executando o projeto
O projeto possui um Makefile com comandos utilitários para execução do projeto listados abaixo:

- Constroi a imagem e sobe o container do projeto build:
```
$ make build-run
```

- Sobe o container do projeto sem build:
```
$ make run
```

## Interagindo com a API
A aplicação possui um tipo de API:
- HTTP REST: responde na porta 8080

### Consultando temperatura por CEP
Após subir o projeto, pode-se consultar a temperatura de uma cidade através de seu CEP da seguinte maneira:

- **Via HTTP:** Execute o arquivo "./api/get_temperature.http".
