# Desafio PosGoExpert - Rate Limiter
Este repositório contém uma implementação de um Rate Limiter em Go. O objetivo deste projeto é limitar o número de requisições por segundo com base em dois critérios: endereço IP e token de acesso. O rate limiter pode ser configurado para aplicar limites tanto para o IP quanto para o token de acesso, e prioriza as configurações do token, caso configurado.

## O rate limiter oferece dois tipos de limitação:
- **Limitação por Endereço IP:** Limita o número de requisições que um determinado endereço IP pode realizar dentro de um intervalo de tempo configurável.
- **Limitação por Token de Acesso:** Limita o número de requisições que um token de acesso pode realizar dentro de um intervalo de tempo configurável. O token é enviado no cabeçalho da requisição com o formato:
```
API_KEY: <TOKEN>
```
**Prioridade de Limitação:** Se o IP e o token de acesso possuem limites configurados, a limitação do token prevalece sobre a do IP.

## Exemplo de Comportamento
- **Limitação por IP:** Se configurado para permitir 5 requisições por segundo por IP e o IP 192.168.1.1 envia 6 requisições em 1 segundo, a sexta requisição será bloqueada.
- **Limitação por Token:** Se o token abc123 estiver configurado com um limite de 10 requisições por segundo e envia 11 requisições, a décima primeira será bloqueada.
Após o tempo de expiração configurado (ex.: 5 minutos), as requisições poderão ser feitas novamente.

## Configuração do Rate Limiter
O rate limiter pode ser configurado de duas formas:
- **Por IP:** Limite de requisições baseado no endereço IP da requisição.
- **Por Token:** Limite de requisições baseado no token de acesso, que será passado no cabeçalho da requisição.

### O rate limiter deve responder com o código HTTP 429 quando o limite de requisições for atingido:
- **Código HTTP:** 429
- **Mensagem:** you have reached the maximum number of requests or actions allowed within a certain time frame

### Exemplo de Arquivo .env
```
MAX_REQUESTS_IP=2
MAX_REQUESTS_TOKEN=2
BLOCK_TIME=10
REDIS_ADDR=localhost:6322
```

## Iniciando o servidor:
```
docker compose up -d
```

## Fazendo uma requisição:
- **Por IP:** curl -X GET http://localhost:8080
- **Por Token:** curl -X GET http://localhost:8080 -H "API_KEY: abc123"