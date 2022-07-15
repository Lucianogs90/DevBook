package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
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

//Buscar traz todos os usuários que atendem a um filtro de nome ou nick
func (respositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick% => os dois primeiros e os dois últimos '%' são carcteres de escape

	linhas, erro := respositorio.db.Query(
		"SELECT id, nome, nick, email, criacao FROM usuarios WHERE nome LIKE ? OR nick LIKE ?",
		nomeOuNick,
		nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criacao,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

//BuscaPorID traz um usuário específico do banco de dados
func (repositorio Usuarios) BuscaPorID(ID uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("SELECT id, nome, nick, email, criacao FROM usuarios WHERE id = ?", ID)
	if erro != nil {
		return modelos.Usuario{}, nil
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.Criacao,
		); erro != nil {
			return modelos.Usuario{}, nil
		}
	}

	return usuario, nil
}
