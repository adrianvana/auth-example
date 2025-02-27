package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken genera un token JWT para el nombre de usuario dado.
// El token contiene el nombre de usuario como un claim y un tiempo de expiración establecido a 24 horas desde el momento de la generación.
// Devuelve el token firmado como una cadena o un error si el proceso de firma falla.
func GenerateToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken valida un token JWT y devuelve si el token es válido.
// Toma un token string como entrada y devuelve un puntero a Claims y un error.
// Si el token es inválido o hay un error durante el análisis, devuelve un error.
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
