package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

//Publicacoes representa um repositório de publicações
type Publicacoes struct {
	DB *sql.DB
}

func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (uint64, error) {
	statement, erro := repositorio.DB.Prepare("INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	publicacaoID, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(publicacaoID), nil
}
