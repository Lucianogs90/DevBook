package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

//Representa um repositório de Usuarios
type Usuarios struct {
	db *sql.DB
}

//NovoRepositorioDeUsuarios cria um repositório de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

//Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint, error) {

	statement, erro := repositorio.db.Prepare(
		"INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	insercao, erro := statement.Exec(
		usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha,
	)
	if erro != nil {
		return 0, erro
	}

	ID, erro := insercao.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint(ID), nil
}
