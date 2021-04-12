# RestServer-Go

RestServer-Go es un servicio RESTFul para administrar tickets y sus respectivos usuarios. Usa dos contenedores uno para el servidor y otro para la base de datos.

## Instalacion


```bash
git clone https://github.com/DamnDanielV/RestServer-Go.git
cd RestServer-Go/
docker-compose up
```



## Uso

Comprobar la conexión al contenedor del servidor

```bash
curl http://localhost:5000/ping
     /// respuesta: pong
```
### Endpoints

#### Tickets
- Obtener todos los tickets de la base de datos:
```
// Metodo: GET
http://localhost:5000/tickets/
```
- Crear un nuevo ticket
```
// Metodo: POST
http://localhost:5000/tickets/
// Body: 
{
    "UserID": 1,
    "status": "abierto"
}
```
- Modificar un ticket
```
// Metodo: PUT
http://localhost:5000/tickets/id
// Body: 
{
    "UserID": 2,
    "status": "cerrado"
}
```
- Eliminar un ticket
```
// Metodo: DELETE
http://localhost:5000/tickets/id
```

### Usuarios
- Obtener todos los usuarios de la base de datos:
```
// Metodo: GET
http://localhost:5000/users/
```
- Crear un nuevo usuario
```
// Metodo: POST
http://localhost:5000/users/
// Body: 
{
    "name": "Daniel V",
    "email": "danielv@gmail.com"
}
```
- Modificar un usuario
```
// Metodo: PUT
http://localhost:5000/users/id
// Body: 
{
    "name": "John Doe"
}
```
- Eliminar un usuario
```
// Metodo: DELETE
http://localhost:5000/users/id
```