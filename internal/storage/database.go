package storage

import (
	"alvytsk/url-shortener/pkg/logger"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	once sync.Once
)

func GetDB() *gorm.DB {
	//Используем singleton (с помощью sync.Once)
	once.Do(func() {
		log := logger.GetLogger()

		var err error
		db, err = gorm.Open(sqlite.Open("url-shortener.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("Ошибка подключения к базе данных", err)
		}

		log.Info("Подключение к базе данных установлено.")

		//Выполняем миграции
		if err := db.AutoMigrate(&Url{}); err != nil {
			log.Fatal("Ошибка выполнения миграции", err)
		}

		log.Info("Миграция выполнена успешно.")
	})
	return db
}