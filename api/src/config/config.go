package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//StringConexao é a string de conexão com o mysql
	StringConexao = ""

	//Porta por onde a API estará rodando
	Porta = 0
)

//Carregar vai inicializar as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		fmt.Print("Não carregou")
		log.Fatal(erro)
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		fmt.Print("Erro na Porta")
		Porta = 9000
	}

	StringConexao = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("BD_USUARIO"),
		os.Getenv("BD_SENHA"),
		os.Getenv("BD_NOME"))

}
