package model

import "testing"

func Test_dist(t *testing.T) {
	type args struct {
		x1 float64
		y1 float64
		x2 float64
		y2 float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dist(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2); got != tt.want {
				t.Errorf("dist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_polarACartesiana(t *testing.T) {
	type args struct {
		pos CoordsPolares
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := polarACartesiana(tt.args.pos)
			if got != tt.want {
				t.Errorf("polarACartesiana() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("polarACartesiana() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestArea(t *testing.T) {
	type args struct {
		pos1 CoordsPolares
		pos2 CoordsPolares
		pos3 CoordsPolares
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Area(tt.args.pos1, tt.args.pos2, tt.args.pos3); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDarPerimetro(t *testing.T) {
	type args struct {
		pos1 CoordsPolares
		pos2 CoordsPolares
		pos3 CoordsPolares
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DarPerimetro(tt.args.pos1, tt.args.pos2, tt.args.pos3); got != tt.want {
				t.Errorf("DarPerimetro() = %v, want %v", got, tt.want)
			}
		})
	}
}
