package routes

import (
	"github.com/faruq/restful-api-ppl/src/handlers"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		events := api.Group("/events")
		{
			events.GET("", handlers.GetAllEvents)
			events.GET("/:id", handlers.GetEventByID)
			events.POST("", handlers.CreateEvent)
			events.PUT("/:id", handlers.UpdateEvent)
			events.DELETE("/:id", handlers.DeleteEvent)
		}
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Events Management API",
			"version": "1.0.0",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})
}
