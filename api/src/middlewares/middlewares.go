package middlewares

import (
	"api/src/autenticacao"
	"api/src/respostas"
	"log"
	"net/http"
)

//Logger imprime informações da requisão no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

//Autentica se o usuario que está fazendo a requisição está autenticado
func Autenticar(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.ERRO(w, http.StatusUnauthorized, erro)
			return
		}
		nextFunc(w, r)
	}
}
