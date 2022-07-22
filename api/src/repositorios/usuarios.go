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

//EditarUsuario atualiza um usuário específico no banco de dados
func (repositorio Usuarios) EditarUsuario(ID uint64, usuario modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID)
	if erro != nil {
		return erro
	}

	return nil
}

//ExcluirUsuario exclui as informações de um usuário específico do banco de dados
func (repositorio Usuarios) ExcluirUsuario(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(ID)
	if erro != nil {
		return erro
	}

	return nil
}

//BuscarPorEmail busca por email um usuário no banco de dados e retorna seu id e senha
func (repositorio Usuarios) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return modelos.Usuario{}, erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(
			&usuario.ID,
			&usuario.Senha,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

//Seguir permite que um usuário siga outro
func (repositorio Usuarios) Seguir(seguidorID, seguidoID uint64) error {
	statement, erro := repositorio.db.Prepare("INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(seguidoID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

//DeixarDeSeguir
func (repositorio Usuarios) DeixarDeSeguir(seguidoID, seguidorID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM seguidores WHERE usuario_id =? AND seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(seguidoID, seguidorID); erro != nil {
		return erro
	}

	return nil
}

//MostrarSeguidores
func (repositorio Usuarios) MostrarSeguidores(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("SELECT u.id, u.nome, u.nick, u.email, u.criacao FROM seguidores s INNER JOIN usuarios u ON u.id = s.seguidor_id WHERE s.usuario_id = ?", usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidores []modelos.Usuario

	for linhas.Next() {
		var seguidor modelos.Usuario
		if erro := linhas.Scan(
			&seguidor.ID,
			&seguidor.Nome,
			&seguidor.Nick,
			&seguidor.Email,
			&seguidor.Criacao,
		); erro != nil {
			return nil, erro
		}
		seguidores = append(seguidores, seguidor)
	}

	return seguidores, nil
}

//MostrarSeguindo
func (repositorio Usuarios) MostrarSeguindo(usuarioID uint64) ([]modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query("SELECT u.id, u.nome, u.nick, u.email, u.criacao FROM seguidores s INNER JOIN usuarios u ON u.id = s.usuario_id WHERE s.seguidor_id = ?", usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidos []modelos.Usuario

	for linhas.Next() {
		var seguido modelos.Usuario
		if erro := linhas.Scan(
			&seguido.ID,
			&seguido.Nome,
			&seguido.Nick,
			&seguido.Email,
			&seguido.Criacao,
		); erro != nil {
			return nil, erro
		}
		seguidos = append(seguidos, seguido)
	}

	return seguidos, nil
}

//BuscarSenha traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("SELECT senha FROM usuarios WHERE id = ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario modelos.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}

	return usuario.Senha, nil
}

//MudarSenha altera a senha de um usuario
func (repositorio Usuarios) MudarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET senha = ? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	_, erro = statement.Exec(senha, usuarioID)
	if erro != nil {
		return erro
	}

	return nil
}
