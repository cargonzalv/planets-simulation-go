package main

import "math"

func dist(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func polarACartesiana(pos CoordsPolares) (float64, float64) {
	return pos.radio * math.Cos(pos.angulo), pos.radio * math.Sin(pos.angulo)
}

func area(pos1 CoordsPolares, pos2 CoordsPolares, pos3 CoordsPolares) float64 {
	x1, y1 := polarACartesiana(pos1)
	x2, y2 := polarACartesiana(pos2)
	x3, y3 := polarACartesiana(pos3)

	return math.Abs((x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)) / 2.0)
}

func darPerimetro(pos1 CoordsPolares, pos2 CoordsPolares, pos3 CoordsPolares) float64 {
	x1, y1 := polarACartesiana(pos1)
	x2, y2 := polarACartesiana(pos2)
	x3, y3 := polarACartesiana(pos3)

	return dist(x1, y1, x2, y2) + dist(x1, y1, x3, y3) + dist(x2, y2, x3, y3)
}
