package routes

import (
	"file_uploading/handlers"

	"github.com/gin-gonic/gin"
)

func AllRoutes(r *gin.Engine) {
	r.GET("/", handlers.HealthCheck)
}
