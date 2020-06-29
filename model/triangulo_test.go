package model

import "testing"

func TestTriangulo_trianguloContieneOrigen(t *testing.T) {
	tests := []struct {
		name string
		t    Triangulo
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.trianguloContieneOrigen(); got != tt.want {
				t.Errorf("Triangulo.trianguloContieneOrigen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangulo_sonColinealesCentro(t *testing.T) {
	tests := []struct {
		name string
		t    Triangulo
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.sonColinealesCentro(); got != tt.want {
				t.Errorf("Triangulo.sonColinealesCentro() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangulo_sonColineales(t *testing.T) {
	tests := []struct {
		name string
		t    Triangulo
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.sonColineales(); got != tt.want {
				t.Errorf("Triangulo.sonColineales() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangulo_area(t *testing.T) {
	tests := []struct {
		name string
		t    Triangulo
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.area(); got != tt.want {
				t.Errorf("Triangulo.area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangulo_darPerimetro(t *testing.T) {
	tests := []struct {
		name string
		t    Triangulo
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.darPerimetro(); got != tt.want {
				t.Errorf("Triangulo.darPerimetro() = %v, want %v", got, tt.want)
			}
		})
	}
}
