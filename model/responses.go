package model

// RespuestaClimaGeneral Estructura de respuesta de consulta de clima general (para los ultimos 10 años por defecto)
type RespuestaClimaGeneral struct {
	Sequias        int `json:"sequias" example:"53"`
	Lluvias        int `json:"lluvias" example:"42"`
	DiaPicoLluvias int `json:"diaPicoLluvias" example:"425"`
	Optimos        int `json:"optimos" example:"2"`
}

// RespuestaClima Estructura de respuesta de consulta del clima del universo en un día particular
type RespuestaClima struct {
	Dia   int    `json:"dia" example:"45"`
	Clima string `json:"clima" example:"soleado"`
}
