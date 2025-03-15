package handlers

import (
	"alvytsk/url-shortener/internal/service"
	"alvytsk/url-shortener/pkg/logger"
	"alvytsk/url-shortener/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

// request структура входящего запроса на сокращение
type createLinkRequest struct {
    URL string `json:"url" binding:"required"`
}

// CreateShortLinkHandler godoc
// @Summary Создание короткой ссылки
// @Description Создаёт короткую ссылку из исходного URL
// @Accept json
// @Produce json
// @Param url body createLinkRequest true "Исходный URL"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /shorten [post]
func CreateShortLinkHandler(c *gin.Context) {
    log := logger.GetLogger()
    var req createLinkRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        log.Error("Ошибка парсинга JSON: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
        return
    }

    // Валидация ссылки
    if err := validation.ValidateURL(req.URL); err != nil {
        log.Error("Ошибка валидации URL: ", err)
        c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный URL"})
        return
    }

    svc := service.NewUrlService()
    link, err := svc.CreateShortLink(req.URL)

    log.Info("Создана короткая ссылка: ", link)
    if err != nil {
        log.Error("Ошибка создания ссылки: ", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать ссылку"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "original_url": link.OriginalURL,
        "short_code":   link.ShortCode,
    })
}

// RedirectHandler godoc
// @Summary Редирект на оригинальный URL
// @Description Перенаправляет на оригинальную ссылку по сокращённому коду
// @Param code path string true "Короткий код ссылки"
// @Success 301
// @Failure 404 {object} map[string]string
// @Router /{short_code} [get]
func RedirectHandler(c *gin.Context) {
    log := logger.GetLogger()
    shortCode := c.Param("code")

    svc := service.NewUrlService()
    link, err := svc.GetOriginalLink(shortCode)
    if err != nil {
        log.Error("Ошибка получения ссылки: ", err)
        c.JSON(http.StatusNotFound, gin.H{"error": "Ссылка не найдена"})
        return
    }

    c.Redirect(http.StatusMovedPermanently, link.OriginalURL)
}
