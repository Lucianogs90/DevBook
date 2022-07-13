package controllers

import (
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//CriarUsuario insere um novo usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		log.Fatal(erro)
	}

	db, erro := banco.Conectar()
	if erro != nil {
		log.Fatal(erro)
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	ID, erro := repositorio.Criar(usuario)
	if erro != nil {
		log.Fatal(erro)
	}

	w.Write([]byte(fmt.Sprintf("ID inserido: %d", ID)))

}

//BuscarUsuarios lista todos os usuários salvos no banco
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {

}

//BuscaUsuario pesquisa um usuário salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// parametros := mux.Vars(r)

	// ID, erro := strconv.ParseUint(parametros["id"], 10, 64)
	// if erro != nil {
	// 	log.Fatal(erro)
	// }

	// db, erro := banco.Conectar()
	// if erro != nil {
	// 	log.Fatal(erro)
	// }
	// defer db.Close()

	// resultado, erro := db.Query("SELECT * FROM usuarios WHERE ID = ?", ID)
	// if erro != nil {
	// 	log.Fatal(erro)
	// }
	// defer resultado.Close()

	// var usuario modelos.Usuario
	// if erro = json.NewEncoder(w).Encode(&usuario); erro != nil {
	// 	log.Fatal(erro)
	// }

}

//AlterarUsuario altera um usuário no banco de dados
func AlterarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Alterando usuário"))
}

//DeletarUsuario excluir um usuário do banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário"))
}
