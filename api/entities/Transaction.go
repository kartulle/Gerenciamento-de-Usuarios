package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct { // se a primeira letra do nome da struct for maiúscula significa q é pública. Caso minúscula, é privada.
	ID         string  `json:"id"` //usamos crase para nomear que quando for json, queremos o id minúsculo.
	Quantia    float64 `json:"quantia"`
	Timestamp  string  `json:"timestamp"`
	Descricao  string  `json:"descricao"`
	SenderId   string  `json:"senderId"`
	ReceiverId string  `json:"receiverId"`
}

func NewTransaction(quantia float64, descricao string, senderId string, receiverId string) *Transaction {
	transaction := Transaction{
		ID:         uuid.New().String(),
		Quantia:    quantia,
		Timestamp:  time.Now().String(),
		Descricao:  descricao,
		SenderId:   senderId,
		ReceiverId: receiverId,
	}
	return &transaction
}
