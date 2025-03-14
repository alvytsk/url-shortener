package service

import (
	"alvytsk/url-shortener/internal/storage"
	"alvytsk/url-shortener/pkg/logger"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"strings"

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

    shortCode := generateShortCode(originalURL)

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
    link := &storage.Url{}
    result := s.db.Where("short_code = ?", shortCode).First(link)
    if result.Error != nil {
        return nil, errors.New("ссылка не найдена")
    }
    return link, nil
}

// generateShortCode генерирует короткий код ссылки
func generateShortCode(url string) string {
    hasher := sha1.New()
    hasher.Write([]byte(url))
    sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
    return strings.TrimRight(sha[:8], "=") // 8 символов короткого кода
}