package signinAndSignup

import (
	"context"
	db "signinsignup/database"
	models "signinsignup/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Ruta de registro
func Signup(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Message: "Datos de registro inválidos",
		})
	}

	// Verificar si el usuario ya existe en la base de datos
	collection := db.Client.Database("gouser").Collection("user")
	result := collection.FindOne(context.TODO(), bson.M{
		"username": user.Username,
	})

	existingUser := &models.User{}
	err := result.Decode(existingUser)
	if err == nil {
		return c.Status(fiber.StatusConflict).JSON(models.Response{
			Message: "El usuario ya existe",
		})
	}

	// Generar hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: "Error al generar la contraseña",
		})
	}

	// Insertar usuario en la base de datos
	_, err = collection.InsertOne(context.TODO(), bson.M{
		"username": user.Username,
		"password": string(hashedPassword),
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: "Error al crear el usuario",
		})
	}

	return c.JSON(models.Response{
		Message: "Usuario creado exitosamente",
	})
}
