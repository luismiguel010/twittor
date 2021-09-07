package models

// RespuestaConsultaRelacion tiene el true o el false que se obtiene de consultar la relación entre 2 usuarios
type RespuestaConsultaRelacion struct {
	Status bool `json:"status"`
}
