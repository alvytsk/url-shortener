package storage

import (
	"alvytsk/url-shortener/pkg/logger"
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
    redisClient *redis.Client
    redisOnce   sync.Once
)

// GetRedisClient возвращает подключение к Redis (singleton)
func GetRedis() *redis.Client {
    redisOnce.Do(func() {
        log := logger.GetLogger()

        redisAddr := "redis:6379" // важно! docker-compose имя сервиса (не localhost)

        redisClient := redis.NewClient(&redis.Options{
            Addr: redisAddr,
            DB:   0,
        })

        if err := redisClient.Ping(context.Background()).Err(); err != nil {
            log.Fatal("Ошибка подключения к Redis:", err)
        }

        log.Info("Подключение к Redis успешно.")
    })

    return redisClient
}