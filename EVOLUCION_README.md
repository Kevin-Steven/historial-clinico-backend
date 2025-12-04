# Servicio de Evolucion y Prescripciones

Endpoints iniciales:

- `POST /evoluciones` -> crear una nueva evolucion/prescripcion para una atencion
- `GET /atenciones/{id}/evoluciones` -> listar evoluciones de una atencion

Variables de entorno:

- `EVOLUCION_DB_DSN` -> DSN MySQL (base `historia_clinica`)
- `EVOLUCION_HTTP_PORT` -> puerto HTTP, por defecto `8085`
- `EVOLUCION_JWT_SECRET` -> secreto JWT (mismo que el de Auth)

Ejemplo payload `POST /evoluciones`:

```json
{
  "id_atencion": 3,
  "fecha_hora": "2025-12-04T17:10:00Z",
  "nota_evolucion": "Paciente con mejoria clinica",
  "farmacoterapia": "Paracetamol 500mg c/8h",
  "indicaciones": "Control ambulatorio",
  "id_profesional": 1,
  "id_usuario": 2
}
```
