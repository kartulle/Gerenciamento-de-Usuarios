package entities

import "github.com/google/uuid"

type User struct { // se a primeira letra do nome da struct for maiúscula significa q é pública. Caso minúscula, é privada.
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Surname       string    `json:"surname"`
	Endereco      string    `json:"endereco"`
	NumeroCelular string    `json:"celular"`
	Contas        []Account `json:"contas"`
}

func NewUser(name string, surname string, endereco string, numeroCelular string, contas []Account) *User {
	user := User{
		ID:            uuid.New().String(),
		Name:          name,
		Surname:       surname,
		Endereco:      endereco,
		NumeroCelular: numeroCelular,
		Contas:        contas,
	}

	return &user
}
