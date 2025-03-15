# 🔗 Url Shortener

Минималистичное веб-приложение для сокращения ссылок, написанное на Go.

## 🛠 Технологии:

- **Go**
- **Gin** — веб-фреймворк
- **SQLite** + **GORM** для базы данных
- **Redis** — кеширование ссылок
- **Logrus** — структурированное логирование
- **Swagger** — документация API
- Docker & docker-compose

## 🗂 Структура проекта:

```
url-shortener
├── cmd
│   └── app
│       └── main.go
├── docs              # Swagger документация
├── internal
│   ├── handlers
│   ├── service
│   └── storage
├── pkg
│   ├── logger
│   └── validation
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── links.db (SQLite DB)
```

## 🚀 Запуск приложения:

### Через Docker-compose (рекомендуется)

```bash
docker-compose up --build
```

Приложение доступно по адресу:

```
http://localhost:8080
```

Swagger-документация доступна по адресу:

```
http://localhost:8080/swagger/index.html
```

## 🔥 API

| Method | Endpoint          | Описание                      |
|--------|-------------------|-------------------------------|
| POST   | `/shorten`        | Создать сокращенную ссылку    |
| GET    | `/{short_code}`   | Переход на оригинальный URL   |
| GET    | `/swagger/*any`   | Swagger-документация          |

## 🛠 Используемые технологии:
- Gin
- GORM + SQLite
- Redis (кэш)
- Logrus
- Validator
- Swagger (Swaggo)

## 📖 Документация API

После запуска документация доступна по адресу:

```
http://localhost:8080/swagger/index.html
```

## 📝 Логирование

Приложение выводит логи в стандартный вывод (stdout) в формате:

```
[INFO] 2025-03-14T22:00:00Z Запускаем сервер на порту :8080
```

## 🛠 Дальнейшие планы развития:

- Покрытие тестами
- Статистика ссылок
- Мониторинг и метрики (Prometheus, Grafana)

