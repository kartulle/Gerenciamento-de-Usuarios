package entities

import (
	"github.com/google/uuid"
)

type Banco struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Endereco string  `json:"endereco"`
	Ramos    []*Ramo `json:"funcao"`
}

func NewBanco(name, endereco string, ramos []*Ramo) *Banco {
	banco := Banco{
		ID:       uuid.New().String(),
		Name:     name,
		Endereco: endereco,
		Ramos:    ramos,
	}
	return &banco
}
