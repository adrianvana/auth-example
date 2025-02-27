package tests

import (
	"auth-example/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDB_Success(t *testing.T) {
	config.InitDB()

	assert.NotNil(t, config.DB, "La base de datos debería estar conectada")
}
