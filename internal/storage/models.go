package storage

import "time"

// Link описывает модель хранения ссылок
type Link struct {
    ID          uint      `gorm:"primaryKey"`
    OriginalURL string    `gorm:"not null;uniqueIndex"` // оригинальная ссылка (уникальная)
    ShortCode   string    `gorm:"not null;uniqueIndex"` // код сокращенной ссылки
    CreatedAt   time.Time // дата создания ссылки
}