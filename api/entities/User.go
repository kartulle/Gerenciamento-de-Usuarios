package entities

import "github.com/google/uuid"

type User struct { // se a primeira letra do nome da struct for maiúscula significa q é pública. Caso minúscula, é privada.
	ID          string `json:"id"` //usamos crase para nomear que quando for json, queremos o id minúsculo.
	Name string `json:"name"`
	Age int `json:"age"`
}

func NewUser() *User {
	user := User{
		ID: uuid.New().String(),
	}
	return &user
}