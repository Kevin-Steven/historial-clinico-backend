# Script de variables de entorno para todos los microservicios

# Configuración común
$env:DB_DSN     = "root:root12345@tcp(localhost:3306)/historia_clinica?parseTime=true&charset=utf8mb4&loc=Local"
$env:JWT_SECRET = "una-clave-secreta-fuerte"  # usa el mismo valor en todos los servicios

# Auth / Usuarios
$env:AUTH_DB_DSN     = $env:DB_DSN
$env:AUTH_HTTP_PORT  = "8082"
$env:AUTH_JWT_SECRET = $env:JWT_SECRET

# Pacientes
$env:PACIENTES_DB_DSN     = $env:DB_DSN
$env:PACIENTES_HTTP_PORT  = "8081"
$env:PACIENTES_JWT_SECRET = $env:JWT_SECRET

# Atenciones Clínicas
$env:ATENCIONES_DB_DSN     = $env:DB_DSN
$env:ATENCIONES_HTTP_PORT  = "8083"
$env:ATENCIONES_JWT_SECRET = $env:JWT_SECRET

# Consulta Externa
$env:CONSULTA_EXTERNA_DB_DSN    = $env:DB_DSN
$env:CONSULTA_EXTERNA_HTTP_PORT = "8084"

# Evolución y Prescripciones
$env:EVOLUCION_DB_DSN     = $env:DB_DSN
$env:EVOLUCION_HTTP_PORT  = "8085"
$env:EVOLUCION_JWT_SECRET = $env:JWT_SECRET

# Signos Vitales
$env:SIGNOS_DB_DSN     = $env:DB_DSN
$env:SIGNOS_HTTP_PORT  = "8086"
$env:SIGNOS_JWT_SECRET = $env:JWT_SECRET

Write-Host "Variables de entorno configuradas para todos los microservicios."
