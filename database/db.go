package db

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB cliente
var Client *mongo.Client

// Función para conectar a MongoDB
func ConnectDB() {
	// Configuración del cliente
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conexión al cliente
	Client, _ = mongo.Connect(context.TODO(), clientOptions)

	// Verificación de conexión
	err := Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conexión a MongoDB exitosa")
}

// Función para cerrar la conexión a MongoDB
func CloseDB() {
	err := Client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Desconexión de MongoDB exitosa")
}
