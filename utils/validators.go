package utils

import (
	"auth-example/models"
	"errors"
	"regexp"
)

// ValidateInput valida los campos de una estructura User.
// Verifica las siguientes condiciones:
// - Los campos Username, Email, Phone y Password no deben estar vacíos.
// - El Email debe tener un formato válido.
// - El Phone debe ser un número de 10 dígitos.
// - La Password debe tener entre 6 y 12 caracteres de longitud.
// - La Password debe contener al menos una letra minúscula, una letra mayúscula, un dígito y un carácter especial de @$&.
// Devuelve un error si alguna de las condiciones no se cumple.
func ValidateInput(user models.User) error {
	if user.Username == "" {
		return errors.New("username es requerido")
	}
	if user.Email == "" {
		return errors.New("email es requerido")
	}
	if user.Phone == "" {
		return errors.New("phone es requerido")
	}
	if user.Password == "" {
		return errors.New("password es requerido")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, user.Email); !matched {
		return errors.New("email inválido")
	}

	phoneRegex := `^\d{10}$`
	if matched, _ := regexp.MatchString(phoneRegex, user.Phone); !matched {
		return errors.New("teléfono inválido, debe tener 10 dígitos")
	}

	// Validar longitud de la contraseña
	if len(user.Password) < 6 || len(user.Password) > 12 {
		return errors.New("contraseña inválida, debe tener entre 6 y 12 caracteres")
	}

	// Validar que contenga al menos una letra minúscula
	if matched, _ := regexp.MatchString(`[a-z]`, user.Password); !matched {
		return errors.New("contraseña inválida, debe contener al menos una letra minúscula")
	}

	// Validar que contenga al menos una letra mayúscula
	if matched, _ := regexp.MatchString(`[A-Z]`, user.Password); !matched {
		return errors.New("contraseña inválida, debe contener al menos una letra mayúscula")
	}

	// Validar que contenga al menos un dígito
	if matched, _ := regexp.MatchString(`[0-9]`, user.Password); !matched {
		return errors.New("contraseña inválida, debe contener al menos un número")
	}

	// Validar que contenga al menos uno de los caracteres especiales @$&
	if matched, _ := regexp.MatchString(`[@$&]`, user.Password); !matched {
		return errors.New("contraseña inválida, debe contener al menos uno de los caracteres especiales @$&")
	}

	return nil
}
