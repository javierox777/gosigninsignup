package main

import (
	"log"
	Signin "signinsignup/controllers"
	db "signinsignup/database"
	models "signinsignup/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Middleware para verificar token JWT
func authMiddleware(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	// Verificar si el token está presente
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Message: "Acceso no autorizado",
		})
	}

	// Verificar token JWT

	return c.Next()
}

// Ruta protegida
func protectedRoute(c *fiber.Ctx) error {
	return c.JSON(models.Response{
		Message: "Ruta protegida",
	})
}

func main() {
	// Conexión a MongoDB
	db.ConnectDB()
	defer db.CloseDB()

	// Configuración de la aplicación Fiber
	app := fiber.New()

	// Middlewares
	app.Use(cors.New())
	app.Use(logger.New())

	// Rutas
	app.Post("/signin", Signin.Signin)
	app.Post("/signup", Signin.Signup)
	app.Get("/protected", authMiddleware, protectedRoute)

	// Iniciar la aplicación en el puerto 3000
	log.Fatal(app.Listen(":3000"))
}
