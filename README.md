# Сервис по укорачиванию ссылок: Ozon

## Информация
- Приложение работает на порту ```8000```
- Использовалось при создании: golang 1.16.6, postgres 13.2, linux debian 10, golangci-lint 1.41.1.</br></br>

## Запуск приложения
```bash
git clone https://github.com/AndreyChemurov/short-cut-link-ozon.git
cd short-cut-link-ozon/
[sudo] docker-compose up
```

## Запуск тестов
```bash
make test
```
или
```bash
go test -v -cover ./...
```

## Запуск линтеров
```bash
make lint
```
или
```bash
golangci-lint run -v ./...
```

## Архитектура

## Результаты выполнения
