package entities

import "github.com/google/uuid"

type Account struct {
	ID          string         `json:"id"`
	NumeroConta string         `json:"numero_conta"`
	Saldo       string         `json:"saldo"`
	TipoConta   string         `json:"tipo_conta"`
	Dono        *User          `json:"usuario"`
	Transacoes  []*Transaction `json:"transacoes"`
}

func NewAccount(numeroConta string, saldo string, tipoConta string, dono *User, transacoes []*Transaction) *Account {
	account := Account{
		ID:          uuid.New().String(),
		NumeroConta: numeroConta, // Set the NumeroConta from the parameter
		Saldo:       saldo,
		TipoConta:   tipoConta,
		Dono:        dono,
		Transacoes:  transacoes,
	}

	return &account
}
