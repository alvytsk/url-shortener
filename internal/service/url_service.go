package service

import (
	"alvytsk/url-shortener/internal/storage"
	"alvytsk/url-shortener/pkg/logger"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"gorm.io/gorm"
)

// UrlService отвечает за бизнес-логику работы со ссылками
type UrlService struct {
	db *gorm.DB
}

// NewUrlService создает новый сервис
func NewUrlService() *UrlService {
    return &UrlService{
        db: storage.GetDB(),
    }
}

// CreateShortLink генерирует и сохраняет сокращённую ссылку
func (s *UrlService) CreateShortLink(originalURL string) (*storage.Url, error) {
    log := logger.GetLogger()

    // Проверка, возможно ссылка уже была сокращена ранее
    existingLink := &storage.Url{}
    result := s.db.Where("original_url = ?", originalURL).First(existingLink)
    if result.Error == nil {
        log.Info("Ссылка уже была сокращена: возвращаем существующую запись")
        return existingLink, nil
    }

    shortCode := generateShortCode(originalURL, 8)
    link := &storage.Url{
        OriginalURL: originalURL,
        ShortCode:   shortCode,
    }

    if err := s.db.Create(link).Error; err != nil {
        log.Error("Ошибка создания ссылки: ", err)
        return nil, err
    }

    return link, nil
}

// GetOriginalLink возвращает оригинальную ссылку по коду
func (s *UrlService) GetOriginalLink(shortCode string) (*storage.Url, error) {
    // log := logger.GetLogger()
    // ctx := context.Background()
    // redisClient := storage.GetRedis()

    // Проверка в Redis
	// cachedURL, err := redisClient.Get(ctx, shortCode).Result()
	// if err == nil {
	// 	log.Info("Получено из Redis cache")
	// 	return &storage.Url{OriginalURL: cachedURL, ShortCode: shortCode}, nil
	// }
    
    link := &storage.Url{}
    result := s.db.Where("short_code = ?", shortCode).First(link)
    if result.Error != nil {
        return nil, errors.New("ссылка не найдена")
    }
    return link, nil
}

// generateShortCode генерирует короткий код ссылки
func generateShortCode(url string, length int) string {
    hash := md5.Sum([]byte(url))
    fullHash := hex.EncodeToString(hash[:])
    if length > len(fullHash) {
        length = len(fullHash)
    }
    return fullHash[:length]
}