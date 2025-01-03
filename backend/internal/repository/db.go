package repository

import (
	"backend/internal/repository/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitDB инициализирует подключение к базе данных
func InitDB(path string) (*gorm.DB, error) {
	// Подключение к базе данных
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}

	// Проверяем, что экземпляр базы данных не равен nil
	if db == nil {
		return nil, fmt.Errorf("gorm.Open вернул nil экземпляр базы данных")
	}

	// Выполнение миграций
	err = migrateDB(db)
	if err != nil {
		return nil, fmt.Errorf("ошибка миграции базы данных: %w", err)
	}

	return db, nil
}

// migrateDB выполняет миграции
func migrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Post{},
	)
}
