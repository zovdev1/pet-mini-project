## запуск backend server

``` bash
go run cmd/main.go
```

## REST API

- /api/v1/user/create
    + POST
    + регистрация пользователь

- /api/v1/user/logIn
    + POST
    + авторизоваться пользователя

- /api/v1/product
    + POST
    + создать товар

- /api/v1/product/:id
    + GET
    + получить товар по айди

- /api/v1/product/?limit=39&offset=0
    + GET
    + получить все товары

## запуск frontend client

```bash
pnpm dev
```

## запуск Docker

```bash
docker-compose up
```