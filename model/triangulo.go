package model

import "math"

// Triangulo Triangulo formado por la posición de los tres planetas
type Triangulo struct {
	punto1 CoordsPolares
	punto2 CoordsPolares
	punto3 CoordsPolares
}

func (t Triangulo) trianguloContieneOrigen() bool {
	origen := CoordsPolares{0, 0}
	//  Calculamos el area del triángulo principal (ABC)
	A := t.area()

	//  Calculamos el área del triángulo (OBC)
	A1 := area(origen, t.punto2, t.punto3)

	//  Calculamos el área del triángulo (OAC)
	A2 := area(origen, t.punto1, t.punto3)

	//  Calculamos el área del triángulo (OAB)
	A3 := area(origen, t.punto1, t.punto2)

	//  Si la suma de los triángulos A1, A2 y A3 resulta en A,
	//  concluimos que el triángulo generado contiene al origen
	return A == A1+A2+A3
}

/*
Calculando si los tres dados puntos son colineales e intersectan al sol
*/
func (t Triangulo) sonColinealesCentro() bool {
	return math.Mod(t.punto1.angulo, 180) == math.Mod(t.punto2.angulo, 180) &&
		math.Mod(t.punto2.angulo, 180) == math.Mod(t.punto3.angulo, 180)
}

/*
Calculando si los tres dados puntos son colineales.
Los 3 puntos son colineales si y solo si el área del triángulo formado es 0
*/
func (t Triangulo) sonColineales() bool {
	return t.area() == 0
}

// Calcula el area usando los 3 puntos de un triángulo
func (t Triangulo) area() float64 {
	return area(t.punto1, t.punto2, t.punto3)
}

// Calcula el perímetro de un triángulo
func (t Triangulo) darPerimetro() float64 {
	return darPerimetro(t.punto1, t.punto2, t.punto3)
}
