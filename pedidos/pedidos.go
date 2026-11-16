package pedidos

import "errors"

// CalcularSubtotal calcula el total de un producto con sus agregados.
func CalcularSubtotal(precioBase float64, preciosAdicionales []float64, cantidad int) float64 {
	var sumaAdicionales float64
	for _, precio := range preciosAdicionales {
		sumaAdicionales += precio
	}
	return (precioBase + sumaAdicionales) * float64(cantidad)
}

// ValidarMonto verifica que el cobro cubra el total de la orden.
func ValidarMonto(totalOrden float64, efectivo float64, transferencia float64) error {
	if (efectivo + transferencia) != totalOrden {
		return errors.New("los montos de pago no coinciden con el total de la orden")
	}
	return nil
}
