# Asteroid API

## Descripción

La API de Asteroides proporciona un conjunto de endpoints para gestionar información sobre asteroides, incluyendo operaciones de creación, lectura, actualización y eliminación. La API está construida utilizando el framework Echo y se conecta a una base de datos MongoDB.

## Requisitos

- Go 1.16+
- MongoDB
- Echo Framework

## Instalación

1. Clona el repositorio:

    ```bash
    git clone https://github.com/AlejandroPintosAlcarazo/asteroid.API.git
    cd asteroid.API
    ```

2. Instala las dependencias:

    ```bash
    go mod tidy
    ```

3. Configura tu base de datos MongoDB y actualiza las variables de entorno según sea necesario.

4. Inicia el servidor:

    ```bash
    go run main.go
    ```

## Endpoints

### Crear un Asteroide

- **URL:** `/api/v1/asteroids`
- **Método:** `POST`
- **Descripción:** Crea un nuevo asteroide.
- **Cuerpo de la solicitud:**

    ```json
    {
        "name": "Asteroid Name",
        "size": "12345",
        "composition": "Type of Composition"
    }
    ```

### Obtener todos los Asteroides

- **URL:** `/api/v1/asteroids`
- **Método:** `GET`
- **Descripción:** Obtiene una lista de todos los asteroides.

### Obtener un Asteroide por ID

- **URL:** `/api/v1/asteroids/:id`
- **Método:** `GET`
- **Descripción:** Obtiene los detalles de un asteroide específico por su ID.
- **Parámetros de URL:**
    - `id` (string): ID del asteroide.

### Actualizar un Asteroide

- **URL:** `/api/v1/asteroids/:id`
- **Método:** `PATCH`
- **Descripción:** Actualiza la información de un asteroide específico.
- **Parámetros de URL:**
    - `id` (string): ID del asteroide.
- **Cuerpo de la solicitud:**

    ```json
    {
        "name": "Updated Name",
        "size": "Updated Size",
        "composition": "Updated Composition"
    }
    ```

### Eliminar un Asteroide

- **URL:** `/api/v1/asteroids/:id`
- **Método:** `DELETE`
- **Descripción:** Elimina un asteroide específico por su ID.
- **Parámetros de URL:**
    - `id` (string): ID del asteroide.

### Eliminar todos los Asteroides

- **URL:** `/api/v1/asteroids`
- **Método:** `DELETE`
- **Descripción:** Elimina todos los asteroides.

### Añadir una Distancia a un Asteroide

- **URL:** `/api/v1/asteroids/:id/distances`
- **Método:** `POST`
- **Descripción:** Añade una nueva distancia a un asteroide específico.
- **Parámetros de URL:**
    - `id` (string): ID del asteroide.
- **Cuerpo de la solicitud:**

    ```json
    {
        "distance": "Distance Value",
        "date": "Date of Measurement"
    }
    ```

### Eliminar una Distancia de un Asteroide

- **URL:** `/api/v1/asteroids/:id/distances/:distanceID`
- **Método:** `DELETE`
- **Descripción:** Elimina una distancia específica de un asteroide.
- **Parámetros de URL:**
    - `id` (string): ID del asteroide.
    - `distanceID` (string): ID de la distancia.

## Middleware

El middleware `serverHeader` añade un encabezado `x-version` con el valor `Test/v1.0` a todas las respuestas del servidor.

## Contribuciones

¡Las contribuciones son bienvenidas! Por favor, abre un issue o envía un pull request.

## Licencia

Este proyecto está bajo la licencia MIT. Consulta el archivo `LICENSE` para más detalles.

## Contacto

Alejandro Pintos Alcarazo - [GitHub](https://github.com/AlejandroPintosAlcarazo)

