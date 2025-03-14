package main

import (
	"alvytsk/url-shortener/internal/handlers"
	"alvytsk/url-shortener/internal/storage"
	"alvytsk/url-shortener/pkg/logger"

	_ "alvytsk/url-shortener/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Link Shortener API
// @version 1.0
// @description API для сокращения и редиректа ссылок.
// @host localhost:8080
// @BasePath /
func main() {
	// Настройка логгера
	log := logger.GetLogger();

	// Подключаемся к базе данных
    db := storage.GetDB()

    // Проверим подключение простым запросом (опционально)
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatal("Ошибка получения объекта sql.DB:", err)
    }

	if err = sqlDB.Ping(); err != nil {
        log.Fatal("Ошибка проверки подключения к БД:", err)
    }

	log.Info("Проверка подключения к БД прошла успешно.")

	// Настройка роутера Gin
	router := gin.Default()

	// Тестовый роут для проверки работы приложения
    router.GET("/ping", func(c *gin.Context) {
        log.Info("Запрос /ping")
        c.JSON(200, gin.H{"message": "pong"})
    })

	router.POST("/shorten", handlers.CreateShortLinkHandler)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/:code", handlers.RedirectHandler)

	log.Info("Запускаем сервер на порту :8080")
    router.Run(":8080")
}