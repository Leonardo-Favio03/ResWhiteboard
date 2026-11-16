package main

import (
	"ResWhiteboard/catalogo"
	"ResWhiteboard/finanzas"
	"ResWhiteboard/pedidos"
	"ResWhiteboard/usuarios"
	"fmt"
)

func main() {
	fmt.Println("=== Iniciando Sistema ===")
	_ = usuarios.RegistrarEmpleado(1, "Cajero 1", "1234")
	_ = catalogo.RegistrarProducto(101, "Combo Clásico", 5.50)

	subtotal := pedidos.CalcularSubtotal(5.50, []float64{1.00}, 1)
	fmt.Printf("Subtotal calculado: $%.2f\n", subtotal)

	_ = finanzas.RegistrarTransaccion(1, 6.50, 0.0)
	efec, transf := finanzas.ObtenerCuadre()
	fmt.Printf("Cuadre actual -> Efectivo: $%.2f | Transferencia: $%.2f\n", efec, transf)
}
