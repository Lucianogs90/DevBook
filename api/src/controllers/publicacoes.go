package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/modelos"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//CriarPublicacao adiciona uma nova publicação ao banco de dados
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioID(r)
	if erro != nil {
		respostas.ERRO(w, http.StatusUnauthorized, erro)
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.ERRO(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacao modelos.Publicacao

	if erro = json.Unmarshal(corpoRequisicao, &publicacao); erro != nil {
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	publicacao.AutorID = usuarioID

	if erro = publicacao.Preparar(); erro != nil{
		respostas.ERRO(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDePublicacoes(db)

	publicacao.ID, erro = repositorio.Criar(publicacao)
	if erro != nil {
		respostas.ERRO(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacao)
}

//BuscarPublicacoes traz todas as publicações próprias e dos usuários que são seguidos
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}

//BuscarPublicacao traz uma publicação específica pelo ID
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}

//AtualizarPublicacao modifica os dados de uma publicação já salva no banco de dados
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}

//DeletarPublicacao exclui os dados de uma determinada publicação já salva no banco de dados
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
