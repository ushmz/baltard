# ratri

[![](https://github.com/ushmz/ratri/actions/workflows/unittest.yml/badge.svg)](https://github.com/ushmz/ratri/actions/workflows/test.yml)

Backend server for my thesis. Works with [savitr](https://github.com/ushmz/savitr)

## Requirements

- Go 1.15+
- Docker

## Run

```shell
docker compose up
```

|service|port|
|:-:|:-:|
|app|8080|
|mysql|3366|
|http(caddy or nginx)|80|
