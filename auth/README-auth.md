# Servicio de Auth/Usuarios

Endpoints iniciales:

- `POST /usuarios` -> crear usuario (ADMIN, MEDICO)
- `POST /auth/login` -> login, devuelve JWT { access_token, rol, usuario_id }

Variables de entorno:

- `AUTH_DB_DSN` -> DSN MySQL (misma base `historia_clinica` o una separada)
- `AUTH_HTTP_PORT` -> puerto HTTP, por defecto `8082`
- `AUTH_JWT_SECRET` -> secreto para firmar JWT

Ejemplo DSN:

```text
AUTH_DB_DSN="root:root12345@tcp(localhost:3306)/historia_clinica?parseTime=true&charset=utf8mb4&loc=Local"
```
