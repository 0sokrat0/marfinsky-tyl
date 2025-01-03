package posts

import (
	"backend/internal/repository/models"

	"gorm.io/gorm"
)

// PostRepository структура репозитория для Post
type PostRepository struct {
	db *gorm.DB
}

// NewRepository создаёт новый экземпляр PostRepository
func NewRepository(db *gorm.DB) Repository {
	return &PostRepository{db: db}
}

// GetAll возвращает все посты
func (r *PostRepository) GetAll() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Find(&posts).Error
	return posts, err
}

// GetByID возвращает пост по ID
func (r *PostRepository) GetByID(id int) (models.Post, error) {
	var post models.Post
	err := r.db.First(&post, id).Error
	return post, err
}

// Create создаёт новый пост
func (r *PostRepository) Create(post models.Post) error {
	return r.db.Create(&post).Error
}

// Update обновляет пост
func (r *PostRepository) Update(post models.Post) error {
	return r.db.Save(&post).Error
}

// Delete удаляет пост по ID
func (r *PostRepository) Delete(id int) error {
	return r.db.Delete(&models.Post{}, id).Error
}
