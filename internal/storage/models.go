package storage

import "time"

// Url описывает модель хранения ссылок
type Url struct {
    ID          uint      `gorm:"primaryKey"`
    OriginalURL string    `gorm:"not null;uniqueIndex"` // оригинальная ссылка (уникальная)
    ShortCode   string    `gorm:"not null;uniqueIndex"` // код сокращенной ссылки
    CreatedAt   time.Time // дата создания ссылки
}