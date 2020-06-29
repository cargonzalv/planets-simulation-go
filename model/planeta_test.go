package model

import "testing"

func TestPlaneta_CalcularDiasPorAnio(t *testing.T) {
	tests := []struct {
		name string
		p    Planeta
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.CalcularDiasPorAnio(); got != tt.want {
				t.Errorf("Planeta.CalcularDiasPorAnio() = %v, want %v", got, tt.want)
			}
		})
	}
}
