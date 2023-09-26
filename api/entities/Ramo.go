package entities

/* import (
	"github.com/google/uuid"
)

type Ramo struct { // se a primeira letra do nome da struct for maiúscula significa q é pública. Caso minúscula, é privada.
	ID               string  `json:"id"` //usamos crase para nomear que quando for json, queremos o id minúsculo.
	Quantia          float64 `json:"quantia"`
	Timestamp        string  `json:"timestamp"` // TODO: Usar timestamp
	Descricao        string  `json:"descricao"`
	ContasEnvolvidas []User  `json:"contas_envolvidas"`
}

func NewRamo() *Ramo {
	user := *NewUser()

	ramo := Ramo{
		ID:               uuid.New().String(),
		Quantia:          0,
		Timestamp:        "",
		Descricao:        "",
		ContasEnvolvidas: []User{user},
	}
	return &ramo
} */
