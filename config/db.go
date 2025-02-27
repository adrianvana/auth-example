package config

import (
	"database/sql"
	"log"
	"sync"

	_ "modernc.org/sqlite"
)

var (
	// DB es el pool de conexiones global de la base de datos.
	DB   *sql.DB
	once sync.Once
)

// InitDB inicializa la conexión a la base de datos y crea las tablas necesarias
// si no existen. Se asegura de que la inicialización se realice solo una vez
// utilizando sync.Once. Si hay un error durante la conexión o la creación de tablas,
// la función registrará el error y terminará la aplicación.
// Patrones de diseño utilizados: Singleton.
func InitDB() {
	once.Do(func() {
		var err error
		DB, err = sql.Open("sqlite", "./app/users.db")
		if err != nil {
			log.Fatal("Error al conectar a la base de datos: ", err)
		}

		_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			phone TEXT UNIQUE NOT NULL,
			password TEXT NOT NULL
		)`)

		if err != nil {
			log.Fatal("Error al crear la tabla: ", err)
		}
	})
}
