package main

import "math"

// Triangulo bla
type Triangulo struct {
	punto1 CoordsPolares
	punto2 CoordsPolares
	punto3 CoordsPolares
}

func (t Triangulo) trianguloContieneOrigen() bool {
	origen := CoordsPolares{0, 0}
	//  Calculate area of triangle ABC
	A := t.area()

	//  Calculate area of triangle OBC
	A1 := area(origen, t.punto2, t.punto3)

	//  Calculate area of triangle OAC
	A2 := area(origen, t.punto1, t.punto3)

	//  Calculate area of triangle OAB
	A3 := area(origen, t.punto1, t.punto2)

	//  Check if sum of A1, A2 and A3
	//  is same as A
	return A == A1+A2+A3
}

func (t Triangulo) sonColinealesCentro() bool {
	return math.Mod(t.punto1.angulo, 180) == math.Mod(t.punto2.angulo, 180) &&
		math.Mod(t.punto2.angulo, 180) == math.Mod(t.punto3.angulo, 180)
}

func (t Triangulo) sonColineales() bool {
	return t.punto1.radio*t.punto2.radio*math.Sin(t.punto2.angulo-t.punto1.angulo)+
		t.punto2.radio*t.punto3.radio*math.Sin(t.punto3.angulo-t.punto2.angulo)+
		t.punto3.radio*t.punto1.radio*math.Sin(t.punto1.angulo-t.punto3.angulo) == 0
}

func (t Triangulo) area() float64 {
	return area(t.punto1, t.punto2, t.punto3)
}

func (t Triangulo) darPerimetro() float64 {
	return darPerimetro(t.punto1, t.punto2, t.punto3)
}
