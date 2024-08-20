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
