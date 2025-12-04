# Servicio de Consulta Externa

Endpoint inicial:

- `POST /consultas-externas` -> crear consulta externa ligada a una atencion

Variables de entorno:

- `CONSULTA_EXT_DB_DSN` -> DSN MySQL (puede ser la misma base `historia_clinica`)
- `CONSULTA_EXT_HTTP_PORT` -> puerto HTTP, por defecto `8084`

Ejemplo payload `POST /consultas-externas`:

```json
{
  "id_atencion": 3,
  "antecedentes_personales": "HTA, DM2",
  "antecedentes_familiares": "Padre con cardiopatia",
  "revision_sistemas": "Sin hallazgos",
  "examen_fisico_regional": "Abdomen blando, depresible",
  "plan_tratamiento": "Analgesia, exámenes complementarios"
}
```
