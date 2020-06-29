package model

import "testing"

func TestCrearPlaneta(t *testing.T) {
	CrearPlaneta("Test1", 35, 300)
	CrearPlaneta("Test2", 459, 2000)
	p := CrearPlaneta("Test3", -30, 500)
	pEsperado := Planeta{
		nombre:    "Test3",
		velocidad: -30,
		posicion: CoordsPolares{
			radio:  500,
			angulo: 0,
		},
	}

	tamanioEsperado := 3
	if len(Planetas) != tamanioEsperado {
		t.Error("Error: Tamaño esperado de ", tamanioEsperado, "tamaño recibido de ", len(Planetas))
	}
	if Planetas[2].nombre != pEsperado.nombre || Planetas[2].velocidad != pEsperado.velocidad ||
		Planetas[2].posicion.angulo != pEsperado.posicion.angulo ||
		Planetas[2].posicion.radio != pEsperado.posicion.radio {
		t.Errorf("Error creando el planeta. Planeta producido: %+v, Planeta esperado:%+v", p, pEsperado)
	}
}
