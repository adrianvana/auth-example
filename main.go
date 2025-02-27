package main

import (
	"auth-example/config"
	"auth-example/routes"
)

// main es la función principal que inicia el servidor.
// Inicializa la base de datos y configura las rutas.
// El servidor se ejecuta en el puerto 8080.
func main() {
	config.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}
