package routes

import (
	"auth-example/controllers"
	"auth-example/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter inicializa y devuelve un nuevo motor Gin con rutas predefinidas.
// Configura las siguientes rutas:
// - POST /register: manejado por controllers.Register
// - POST /login: manejado por controllers.Login
// Devuelve un puntero a la instancia configurada de gin.Engine.

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middlewares.ValidateToken())
	{
		protected.GET("/random-user", controllers.GetRandomUser)
	}

	return r
}
