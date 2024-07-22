# Desafio Prático Sistema de Temperatura por CEP
Projeto do Desafio Prático para conclusão da Pós Graduação em Go Expert da Full Cycle.

<p align="center">
  <img src="https://blog.golang.org/gopher/gopher.png" alt="">
</p>

Esta é uma aplicação desenvolvida em Go, cujo objetivo é receber um CEP, validar se possui 8 caracteres, consultar a API ViaCEP para obter a localização, consultar a Weather API e obter a temperatura em Celsius. Por fim esta API retorna a temperatura em Celsius, Fahrenheit e Kelvin.

<br>

## Índice


- [Instalação](#instalação)
- [Pré requisitos](#pré-requisitos)
- [Como Usar](#como-usar)
- [Testes unitários e de Integração](#testes-unitários-e-de-integração)
- [Google Cloud Run](#google-cloud-run)
- [Contato](#contato)
- [Agradecimentos](#agradecimentos)

<br>


## Instalação

```sh
$ git clone https://github.com/gilbertom/go-temperatura-cep.git
$ docker-compose up --build
```
<br>

## Pré requisitos
Esta aplicação utiliza a WeatherAPI e é obrigatório ter uma chave de acesso.  
Necessário também ter o Go e Docker instalado.

1. Crie uma conta no site <a href="https://www.weatherapi.com">WeatherAPI</a>.
2. Na página <a href="https://www.weatherapi.com/my/">Home</a>, copie a chave de acesso (API Key).
3. Adicione a chave de acesso no arquivo ".env" localizado na pasta /cmd/app, utilizando a variável API_KEY_WEATHER.

<br>


## Como Usar

No diretório /api temos dois arquivos. 

  **get_temperature.http**  
  Realiza uma chamada GET no localhost:8080

  **get_temperature_cloud.http**  
  Realiza uma chamada GET no serviço hospedado no Google Cloud Run.

Obs.: Imprescindível instalar a Extensão 'HTTP Client' no seu Visual Studio Code.  

Request de um CEP válido
```sh
  GET http://localhost:8080?cep=28951620 HTTP/1.1
  Host: localhost:8080
  Content-Type: application/json
```

Response
  ```sh
  HTTP/1.1 200 OK
  Content-Type: application/json
  Date: Mon, 22 Jul 2024 16:50:36 GMT
  Content-Length: 58
  Connection: close

  {
    "temp_C": 29.3,
    "temp_F": 84.74000000000001,
    "temp_K": 302.3
  }
  ```
<br>  

___

  Request de um CEP inválido
  ```sh
    GET http://localhost:8080?cep=2895162A HTTP/1.1
    Host: localhost:8080
    Content-Type: application/json
  ```

  Response
  ```sh
  HTTP/1.1 422 Unprocessable Entity
  Content-Type: text/plain; charset=utf-8
  X-Content-Type-Options: nosniff
  Date: Mon, 22 Jul 2024 16:51:47 GMT
  Content-Length: 16
  Connection: close

  invalid zipcode
  ```
<br>  

___

  Request de um CEP inexistente
  ```sh
  GET http://localhost:8080?cep=88888888 HTTP/1.1
  Host: localhost:8080
  Content-Type: application/json
  ```

  Response
  ```sh
  HTTP/1.1 404 Not Found
  Content-Type: text/plain; charset=utf-8
  X-Content-Type-Options: nosniff
  Date: Mon, 22 Jul 2024 16:58:17 GMT
  Content-Length: 21
  Connection: close

  can not find zipcode
  ```
<br>  


## Testes unitários e de Integração
Para executar os testes unitários e de Integração, execute o comando:
```sh
$ go test ./...
```

<br>

## Google Cloud Run
Esta aplicação está hospedada na plataforma Google Cloud Run. Você pode acessá-la através do link: <a href="https://clean-weather-vccdavtg5q-uc.a.run.app/?cep=28951620">API_WEATHER</a> ou diretamente pelo endereço: https://clean-weather-vccdavtg5q-uc.a.run.app/?cep=28951620.

<br>

## Contato
Para entrar em contato com o desenvolvedor deste projeto:
[gilbertomakiyama@gmail.com](mailto:gilbertomakiyama@gmail.com)

<br>




## Agradecimentos
Gostaria de expressar minha sincera gratidão a todo o time do curso de Pós-Graduação em Go Avançado da FullCycle pelo empenho, dedicação e excelência no ensino. Suas contribuições foram fundamentais para o meu desenvolvimento e sucesso. Muito obrigado!
