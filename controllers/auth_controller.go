package controllers

import (
	"auth-example/config"
	"auth-example/models"
	"auth-example/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register maneja el proceso de registro de usuarios.
// Espera una carga útil JSON con los detalles del usuario, valida la entrada,
// verifica si el usuario ya existe, hashea la contraseña e inserta
// el nuevo usuario en la base de datos.
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	if err := utils.ValidateInput(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var exists int
	config.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = ? OR phone = ?", user.Email, user.Phone).Scan(&exists)
	if exists > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El correo o teléfono ya está registrado"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	_, err := config.DB.Exec("INSERT INTO users (username, email, phone, password) VALUES (?, ?, ?, ?)", user.Username, user.Email, user.Phone, string(hashedPassword))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al registrar usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario registrado exitosamente"})
}

// Login maneja la autenticación de usuarios verificando el correo/usuario y la contraseña proporcionados.
// Si las credenciales son válidas, genera y devuelve un token JWT.
func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parámetros inválidos"})
		return
	}

	if input.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falta el campo correo"})
		return
	}

	if input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Falta el campo contraseña"})
		return
	}

	var user models.User
	err := config.DB.QueryRow("SELECT id, username, email, phone, password FROM users WHERE email = ? OR username = ?", input.Email, input.Email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Phone, &user.Password)

	if err == sql.ErrNoRows || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario/contraseña incorrectos"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario/contraseña incorrectos"})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
