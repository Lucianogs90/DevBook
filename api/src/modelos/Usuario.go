package modelos

import (
	"time"
)

//Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID      uint      `json:"id,omitempty"`
	Nome    string    `json:"nome,omitempty"`
	Nick    string    `json:"nick,omitempty"`
	Email   string    `json:"email,omitempty"`
	Senha   string    `json:"senha,omitempty"`
	Criacao time.Time `json:"criacao,omitempty"`
}
