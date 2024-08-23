# GO-GIT-SCANNER

Este proyecto es un escáner de Git que realiza diferencias periódicas y las envía a un endpoint.

## Estructura del Proyecto

```
GO-GIT-SCANNER/
├── src/
│   └── utils/
│       ├── diff.go
│       ├── file.go
│       ├── post.go
│       └── scanner.go
├── Makefile
└── README.md
```

## Uso

1. Compilar el proyecto:

   ```
   make build
   ```

2. Ejecutar el escáner:

   ```
   make run
   ```

3. Limpiar los binarios:
   ```
   make clean
   ```

Asegúrate de configurar las constantes en los archivos (`gitDir`, `outputFile`, `apiEndpoint`) antes de ejecutar.

## Dependencias

Este proyecto utiliza:

- github.com/robfig/cron/v3

Asegúrate de ejecutar `go mod tidy` para gestionar las dependencias.

Hola mundo - 2024-08-22 21:33:00

Hola mundo - 2024-08-23 08:29:53

Hola mundo - 2024-08-23 08:34:44

Hola mundo - 2024-08-23 08:53:48

Hola mundo - 2024-08-23 08:57:58

Hola mundo - 2024-08-23 08:58:33

Hola mundo - 2024-08-23 09:01:39

Hola mundo - 2024-08-23 09:03:12

Hola mundo - 2024-08-23 09:08:03

Hola mundo - 2024-08-23 09:10:55

Hola mundo - 2024-08-23 09:27:06

Hola mundo - 2024-08-23 09:28:27

Hola mundo - 2024-08-23 09:36:51

Hola mundo - 2024-08-23 09:38:32

Hola mundo - 2024-08-23 09:39:06

Hola mundo - 2024-08-23 09:40:46

Hola mundo - 2024-08-23 16:07:41

Hola mundo - 2024-08-23 16:18:43

Hola mundo - 2024-08-23 16:20:15

Hola mundo - 2024-08-23 16:21:26

Hola mundo - 2024-08-23 16:27:49
