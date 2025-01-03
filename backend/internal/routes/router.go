package routes

import (
	"backend/internal/core/posts"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes инициализирует маршруты
func SetupRoutes(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // URL вашего фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Настройка маршрутов для постов
	postHandler := posts.NewHandler(posts.NewService(posts.NewRepository(db)))

	api := router.Group("/api")
	{
		postsGroup := api.Group("/posts")
		{
			postsGroup.GET("", postHandler.GetAll)     // GET /api/posts
			postsGroup.HEAD("", func(c *gin.Context) { // HEAD /api/posts
				c.Status(200)
			})
			postsGroup.POST("", postHandler.Create)       // POST /api/posts
			postsGroup.GET("/:id", postHandler.GetByID)   // GET /api/posts/:id
			postsGroup.PUT("/:id", postHandler.Update)    // PUT /api/posts/:id
			postsGroup.DELETE("/:id", postHandler.Delete) // DELETE /api/posts/:id
		}
	}

	return router
}
