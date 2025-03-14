.PHONY: swag run

# Генерация документации
swag:
	swag init -g cmd/app/main.go

# Запуск приложения
run: swag
	go run cmd/app/main.go