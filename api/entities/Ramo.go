package entities

import (
	"github.com/google/uuid"
)

type Ramo struct { // se a primeira letra do nome da struct for maiúscula significa q é pública. Caso minúscula, é privada.
	ID         string      `json:"id"`
	Nome       string      `json:"nome"` // Updated to string
	Endereco   string      `json:"endereco"`
	Contas     []User      `json:"contas"`
	Empregados []Empregado `json:"empregados"`
}

func NewRamo(nome, endereco string, contas []User, empregados []Empregado) *Ramo {
	ramo := Ramo{
		ID:         uuid.New().String(),
		Nome:       nome,
		Endereco:   endereco,
		Contas:     contas,
		Empregados: empregados,
	}
	return &ramo
}
