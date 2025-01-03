package posts

import (
	"backend/internal/repository/models"
	"backend/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler структура для хендлеров
type Handler struct {
	service *Service
}

// NewHandler создаёт новый экземпляр Handler
func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// GetAll обрабатывает запрос на получение всех постов
func (h *Handler) GetAll(c *gin.Context) {
	posts, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error(), "Не удалось получить посты"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(posts, "Посты успешно получены"))
}

// GetByID обрабатывает запрос на получение поста по ID
func (h *Handler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), "Некорректный ID поста"))
		return
	}

	post, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse(err.Error(), "Пост не найден"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(post, "Пост успешно получен"))
}

// Create обрабатывает запрос на создание нового поста
func (h *Handler) Create(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), "Некорректный запрос"))
		return
	}

	if err := h.service.Create(post); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error(), "Не удалось создать пост"))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(nil, "Пост успешно создан"))
}

// Update обрабатывает запрос на обновление поста
func (h *Handler) Update(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), "Некорректный запрос"))
		return
	}

	if err := h.service.Update(post); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error(), "Не удалось обновить пост"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(nil, "Пост успешно обновлён"))
}

// Delete обрабатывает запрос на удаление поста
func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error(), "Некорректный ID поста"))
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(err.Error(), "Не удалось удалить пост"))
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse(nil, "Пост успешно удалён"))
}
