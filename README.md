# Auth Example

Este proyecto es una aplicación de autenticación simple implementada en Go, que utiliza JWT para la autenticación de usuarios. El proyecto incluye operaciones de registro y login de usuarios, con almacenamiento en una base de datos SQLite.

## Estructura del Proyecto

La estructura del proyecto es la siguiente:

```
/auth-example
│── main.go
│── config
│   └── db.go  // Configuración y conexión a la base de datos
│── controllers
│   ├── auth_controller.go  // Controladores para autenticación
│── models
│   ├── user.go  // Definición del modelo de usuario
│── routes
│   ├── routes.go  // Definición de rutas
│── utils
│   ├── jwt.go  // Funciones para generar y validar JWT
│   ├── validators.go  // Validaciones de entrada como email, teléfono y contraseña
│── Dockerfile
│── docker-compose.yml
```

## Requisitos

- Docker y Docker Compose (opcional, para ejecutar el proyecto y pruebas en contenedores)

- Go 1.21 o superior (si prefieres ejecutar la aplicación localmente sin Docker)

## Instalación

### Clonar el repositorio

```bash
git  clone  https://github.com/tu-usuario/auth-example.git
cd  auth-example
```

### Construir y ejecutar con Docker

Si prefieres ejecutar el proyecto en un contenedor Docker, puedes usar el siguiente comando para construir la imagen y ejecutar la aplicación:

```bash
docker-compose  up  --build
```

Este comando levantará los servicios de la aplicación y la base de datos en contenedores Docker. La aplicación estará disponible en http://localhost:8080.

### Ejecutar pruebas con Docker

Para ejecutar las pruebas unitarias dentro del contenedor Docker, usa el siguiente comando:

```bash
docker-compose run --rm tests
```

Esto ejecutará las pruebas y mostrará el resultado en la consola.

### Ejecutar el proyecto sin Docker

Si prefieres ejecutar el proyecto directamente en tu máquina sin Docker, sigue estos pasos:

- Instalar las dependencias de Go: 

```bash
go mod download
```

- Iniciar la aplicación:

```bash
go run main.go
```

La aplicación estará disponible en http://localhost:8080.

### Endpoints

- POST /register

Este endpoint permite registrar un nuevo usuario. Los parámetros de entrada deben ser enviados en formato JSON:

Body:
```json
{
"username": "usuario",
"email": "usuario@ejemplo.com",
"phone": "1234567890",
"password": "contraseña123"
}
```

Respuesta exitosa (200):
```json
{
"message": "Usuario registrado exitosamente"
}
```

Respuesta de error (400):
```json
{
"error": "El correo o teléfono ya está registrado"
}
```

-POST /login

Este endpoint permite que los usuarios inicien sesión utilizando su nombre de usuario o correo electrónico y su contraseña. Los parámetros deben enviarse en formato JSON:

Body:
```json
{
"emailOrUsername": "usuario",
"password": "contraseña123"
}
```

Respuesta exitosa (200):
```json
{
"token": "jwt_token_aqui"
}
```

Respuesta de error (401):
```json
{
"error": "Usuario/contraseña incorrectos"
}
```

### Pruebas

El proyecto incluye pruebas unitarias para los controladores y la configuración de la base de datos.
Para ejecutar las pruebas, puedes usar uno de los siguientes métodos:

- Ejecutar pruebas dentro de Docker:

```bash
docker-compose run --rm tests
```

- Ejecutar pruebas localmente (sin Docker):

```bash
go test ./tests
```

- Coverage

Puedes obtener el coverage de las pruebas de la siguiente manera:

- Ejecuta las pruebas con cobertura:

```bash
go test -cover ./tests
```

- Generar un informe detallado en HTML:

```bash
go test -coverprofile=coverage.out

go tool cover -html=coverage.out -o coverage.html
```

Abre el archivo coverage.html en tu navegador para ver un reporte visual.

### Tecnologías

Go: Lenguaje de programación principal.

Gin: Framework web para Go.

SQLite: Base de datos ligera utilizada para almacenar los datos de los usuarios.

JWT: Para la autenticación y autorización de usuarios.

Contribuciones

Las contribuciones son bienvenidas. Si tienes alguna sugerencia o mejora, no dudes en abrir un issue o enviar un pull request.

### Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo LICENSE para más detalles.

### Explicación:

  
- **Estructura**: El README explica la estructura del proyecto y los archivos importantes.

- **Instrucciones de instalación**: Incluye cómo clonar el repositorio, ejecutar el proyecto con Docker y sin Docker, y ejecutar las pruebas.

- **Endpoints**: Detalla los endpoints disponibles en la API (registro y login).

- **Pruebas**: Explica cómo ejecutar pruebas y obtener cobertura de pruebas.

- **Tecnologías**: Enumera las tecnologías utilizadas en el proyecto.

- **Contribuciones**: Incluye una sección de contribuciones para quienes deseen participar en el proyecto.

- **Licencia**: La licencia MIT para que otros puedan utilizar el proyecto.

  