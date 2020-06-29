package model

import (
	"reflect"
	"testing"
)

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

func TestCrearPlanetas(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CrearPlanetas()
		})
	}
}

func Test_agregarPlaneta(t *testing.T) {
	type args struct {
		p Planeta
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			agregarPlaneta(tt.args.p)
		})
	}
}

func TestBuscarPlanetaPorNombre(t *testing.T) {
	type args struct {
		nombre string
	}
	tests := []struct {
		name string
		args args
		want *Planeta
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuscarPlanetaPorNombre(tt.args.nombre); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuscarPlanetaPorNombre() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalcularClimaDia(t *testing.T) {
	type args struct {
		dia int
	}
	tests := []struct {
		name string
		args args
		want RespuestaClima
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcularClimaDia(tt.args.dia); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalcularClimaDia() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_darEstado(t *testing.T) {
	type args struct {
		t Triangulo
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := darEstado(tt.args.t); got != tt.want {
				t.Errorf("darEstado() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSimulacion(t *testing.T) {
	type args struct {
		dias int
	}
	tests := []struct {
		name string
		args args
		want RegistroClima
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Simulacion(tt.args.dias, false); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Simulacion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_avanzarDias(t *testing.T) {
	type args struct {
		dia int
	}
	tests := []struct {
		name string
		args args
		want Triangulo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := avanzarDias(tt.args.dia); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avanzarDias() = %v, want %v", got, tt.want)
			}
		})
	}
}
