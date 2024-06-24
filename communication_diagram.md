# Communication Diagram

```bash
+-------------------+
|       Usuario     |
+-------------------+
          |
          v
+-------------------+
|   Control-Panel   |
|   Adapter Driver  |
+-------------------+
          |
          v
+-----------------------------+
|       Control-Panel         |
|      Adapter Driven         |
|                             |
|  +---------------------+    |
|  | Control-Panel Port  |    |
|  +---------------------+    |
|                             |
|   +-------------------+     |
|   | YouTube Port      |     |
|   | Monitor Port      |     |
|   | Logs Port         |     |
|   +-------------------+     |
+-----------------------------+
          |
          v
+-----------------------------+
|        YouTube-API          |
|      Adapter Driver         |
+-----------------------------+
          |
          v
+-----------------------------+
|        YouTube-API          |
|       Adapter Driven        |
|                             |
|  +---------------------+    |
|  | YouTube Port        |    |
|  +---------------------+    |
+-----------------------------+
          |
          v
+-----------------------------+
|      Monitor-API            |
|      Adapter Driver         |
+-----------------------------+
          |
          v
+-----------------------------+
|       Monitor-API           |
|       Adapter Driven        |
|                             |
|  +---------------------+    |
|  | Monitor Port        |    |
|  +---------------------+    |
+-----------------------------+
          |
          v
+-----------------------------+
|        Logs-API             |
|      Adapter Driver         |
+-----------------------------+
          |
          v
+-----------------------------+
|         Logs-API            |
|       Adapter Driven        |
|                             |
|  +---------------------+    |
|  | Logs Port           |    |
|  +---------------------+    |
+-----------------------------+
          |
          v
+-----------------------------+
|      Control-Panel          |
|    Adapter Driver           |
+-----------------------------+
          |
          v
+-------------------+
|      Usuario      |
+-------------------+
```

## Diagram Explanation

Usuario: Envía una solicitud al Control-Panel.

Control-Panel Adapter Driver: Recibe la solicitud y la pasa al Control-Panel Adapter Driven.

Control-Panel Adapter Driven: Se comunica con el YouTube-API para obtener la información solicitada. También se comunica con los hexágonos de Monitor-API y Logs-API para registrar la actividad y manejar los errores.

YouTube-API Adapter Driver: Se comunica con el YouTube-API Adapter Driven para obtener la información de YouTube.

YouTube-API Adapter Driven: Usa el puerto YouTube Port para obtener la información de YouTube y la devuelve al Control-Panel Adapter Driven.

Monitor-API Adapter Driver: Se comunica con el Monitor-API Adapter Driven para registrar la actividad del usuario.

Monitor-API Adapter Driven: Usa el puerto Monitor Port para guardar la información de monitoreo.

Logs-API Adapter Driver: Se comunica con el Logs-API Adapter Driven para manejar los errores.

Logs-API Adapter Driven: Usa el puerto Logs Port para guardar la información de errores.

Control-Panel Adapter Driver: Recibe la información de todos los hexágonos y la devuelve al usuario en formato JSON.