package tests

import (
	"auth-example/config"
	"auth-example/controllers"
	"auth-example/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

// setupRouter inicializa un nuevo motor Gin con rutas predefinidas.
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	return r
}

// TestRegister_Success prueba la funcionalidad de registro para un intento exitoso.
func TestRegister_Success(t *testing.T) {
	config.InitDB()
	mockUser := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Phone:    "1234567890",
		Password: "Abc123$",
	}

	mockUserJSON, _ := json.Marshal(mockUser)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(mockUserJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Usuario registrado exitosamente")
}

// TestRegister_DuplicateUser prueba la funcionalidad de registro para una contraseña no válida.
func TestRegister_InvalidPassword(t *testing.T) {
	config.InitDB()
	mockUser := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Phone:    "1234567890",
		Password: "password123",
	}

	mockUserJSON, _ := json.Marshal(mockUser)

	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(mockUserJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "contraseña inválida")
}

// TestRegister_InvalidInput prueba la funcionalidad de registro para una carga útil JSON inválida.
func TestRegister_InvalidInput(t *testing.T) {
	config.InitDB()
	mockUserJSON := `{"username":"testuser","email":"test@example.com","phone":"1234567890}`
	req, _ := http.NewRequest("POST", "/register", bytes.NewBufferString(mockUserJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "Parámetros inválidos")
}

// TestLogin_Success prueba la funcionalidad de inicio de sesión para un intento exitoso.
func TestLogin_Success(t *testing.T) {
	config.InitDB()
	mockUser := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Phone:    "1234567890",
		Password: "Abc123$",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(mockUser.Password), bcrypt.DefaultCost)
	config.DB.Exec("INSERT INTO users (username, email, phone, password) VALUES (?, ?, ?, ?)",
		mockUser.Username, mockUser.Email, mockUser.Phone, string(hashedPassword))

	loginJSON := `{"emailOrUsername":"testuser","password":"password123"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

func TestLogin_Fail_InvalidCredentials(t *testing.T) {
	config.InitDB()
	loginJSON := `{"emailOrUsername":"wronguser","password":"wrongpassword"}`
	req, _ := http.NewRequest("POST", "/login", bytes.NewBufferString(loginJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Usuario/contraseña incorrectos")
}
