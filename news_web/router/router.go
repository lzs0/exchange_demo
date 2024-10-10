package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"news_web/controllers"
	"news_web/middlewares"
	"time"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleWare())
	{
		api.GET("/exchangeRates", controllers.GetExchangeRate)
		api.POST("/exchangeRates", controllers.CreateExchangeRate)

		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.GetArticles)
		api.GET("articles/:id", controllers.GetArticleById)

		api.POST("/articles/:id/like", controllers.LikeArticle)
		api.GET("/articles/:id/like", controllers.GetArticleLikes)
	}

	return r
}
