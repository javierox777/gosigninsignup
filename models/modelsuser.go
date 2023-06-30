package modeluser

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Estructura de respuesta
type Response struct {
	Message string `json:"message"`
}
