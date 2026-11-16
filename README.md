# 🍔 ResWhiteboard  (Go)

Proyecto desarrollado para la asignatura de **Programación Orientada a Objetos (Go)**. Esta diseñado para personal de mostrador en locales de comida, busca facilitar y agilizar el proceso de toma de pedidos y entrega de facturas.

## 🎯 Objetivo del Proyecto
Agilizar la toma de pedidos con adicionales y automatizar el cuadre de caja diario, garantizando una separación estricta e inmutable entre los ingresos recibidos en **Efectivo** y por **Transferencia Bancaria**, asegurando la trazabilidad de cada empleado que opera el sistema.

## 🏗️ Arquitectura y Paquetes (Módulos del Sistema)
El proyecto sigue un enfoque modular basado en funciones puras y paquetes independientes en minúsculas, omitiendo estructuras (`structs`) en esta primera etapa:

* **`usuarios/`**: Gestiona el registro de empleados y la validación de identidad mediante códigos PIN.
* **`catalogo/`**: Administra los productos base del menú y los adicionales con su respectivo costo.
* **`pedidos/`**: Contiene el motor matemático para el cálculo de subtotales con adicionales y validación de pagos mixtos.
* **`finanzas/`**: Centraliza el registro histórico de transacciones y ejecuta el desglose del cuadre de caja separado por canal de pago.

## 📂 Estructura de Directorios

```text
ResWhiteboard/
├── go.mod                   # Definición del módulo único de Go
├──.gitignore                # ignorar esto(paqute de preferencias del ide)
├── README.md                # Documentación del proyecto
├── main.go                  # Punto de entrada y orquestación
├── usuarios/
│   └── usuarios.go          # Funciones de acceso y auditoría
├── catalogo/
│   └── catalogo.go          # Funciones de menú y precios
├── pedidos/
│   └── pedidos.go           # Funciones de cálculo de órdenes
└── finanzas/
    └── finanzas.go          # Funciones de control de caja
```

## 🚀 Cómo verificar y compilar el proyecto

Asegúrate de tener instalado **Go (versión 1.21 o superior)**. Clona este repositorio y ejecuta los siguientes comandos en la raíz del proyecto:

1. **Verificar la compilación de todos los paquetes:**
   ```bash
   go build ./...
   ```
   *(Si el comando finaliza sin mostrar mensajes en consola, todos los paquetes compilan correctamente).*

2. **Ejecutar la demostración básica:**
   ```bash
   go run main.go
   ```
