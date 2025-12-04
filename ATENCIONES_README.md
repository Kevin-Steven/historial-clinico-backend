# Servicio de Atenciones Clinicas

Endpoints iniciales:

- `POST /atenciones` -> crear una nueva atencion (encounter)

Variables de entorno:

- `ATENCIONES_DB_DSN` -> DSN MySQL (puede ser la misma base `historia_clinica`)
- `ATENCIONES_HTTP_PORT` -> puerto HTTP, por defecto `8083`

Ejemplo payload `POST /atenciones`:

```json
{
  "id_paciente": 1,
  "id_establecimiento": 1,
  "id_tipo_atencion": 1,
  "motivo_consulta": "Dolor abdominal",
  "enfermedad_actual": "Dolor de 3 dias de evolucion",
  "id_usuario_crea": 2
}
```
