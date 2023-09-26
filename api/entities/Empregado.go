package entities

import (
	"github.com/google/uuid"
)

type Empregado struct { // se a primeira letra do nome da struct for maiúscula significa q é pública. Caso minúscula, é privada.
	ID       string `json:"id"` //usamos crase para nomear que quando for json, queremos o id minúsculo.
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Endereco string `json:"endereco"`
	Celular  string `json:"celular"`
	Funcao   string `json:"funcao"`
}

func NewEmpregado() *Empregado {
	Empregado := Empregado{
		ID:       uuid.New().String(),
		Name:     "",
		Surname:  "",
		Endereco: "",
		Celular:  "",
		Funcao:   "",
	}
	return &Empregado
}
