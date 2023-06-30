package signinAndSignup

import (
	"context"
	db "signinsignup/database"
	modeluser "signinsignup/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Ruta de inicio de sesión
func Signin(c *fiber.Ctx) error {
	user := new(modeluser.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(modeluser.Response{
			Message: "Datos de inicio de sesión inválidos",
		})
	}

	// Buscar usuario en la base de datos
	collection := db.Client.Database("gouser").Collection("user")
	result := collection.FindOne(context.TODO(), bson.M{
		"username": user.Username,
	})

	storedUser := &modeluser.User{}
	err := result.Decode(storedUser)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(modeluser.Response{
			Message: "Credenciales inválidas",
		})
	}

	// Verificar la contraseña
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(modeluser.Response{
			Message: "Credenciales inválidas",
		})
	}

	// Generar token JWT

	return c.JSON(modeluser.Response{
		Message: "Inicio de sesión exitoso",
	})
}
