package modelos

import (
	"errors"
	"strings"
	"time"
)

//Publicacao representa uma publicação feita pelo usuário
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autor_id,omitempty"`
	AutorNick string    `json:"autor_nick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	Criacao   time.Time `json:"criacao,omitempty"`
}

//Preparar vai chamar os métodos para validar e formatar a publicaçao recebida
func (publicacao *Publicacao) Preparar() error {
	if erro := publicacao.validar(); erro != nil {
		return erro
	}

	publicacao.formatar()

	return nil
}

func (publicacao *Publicacao) validar() error {
	if publicacao.Titulo == "" {
		return errors.New("o título é obrigatório e não pode estar em branco")
	}

	if publicacao.Conteudo == "" {
		return errors.New("o conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (publicacao *Publicacao) formatar() {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
}
