package posts

import (
	"backend/internal/repository/models"
	"errors"
)

// Repository интерфейс для работы с репозиторием
type Repository interface {
	GetAll() ([]models.Post, error)
	GetByID(id int) (models.Post, error)
	Create(post models.Post) error
	Update(post models.Post) error
	Delete(id int) error
}

// Service структура для бизнес-логики
type Service struct {
	repo Repository
}

// NewService создаёт новый экземпляр Service
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// GetAll возвращает все посты
func (s *Service) GetAll() ([]models.Post, error) {
	return s.repo.GetAll()
}

// GetByID возвращает пост по ID
func (s *Service) GetByID(id int) (models.Post, error) {
	return s.repo.GetByID(id)
}

// Create создаёт новый пост
func (s *Service) Create(post models.Post) error {
	if post.Title == "" || post.Body == "" {
		return errors.New("заголовок и содержимое поста не могут быть пустыми")
	}
	return s.repo.Create(post)
}

// Update обновляет существующий пост
func (s *Service) Update(post models.Post) error {
	if post.ID == 0 {
		return errors.New("ID поста обязателен для обновления")
	}
	return s.repo.Update(post)
}

// Delete удаляет пост по ID
func (s *Service) Delete(id int) error {
	return s.repo.Delete(id)
}
