package seguranca

import "golang.org/x/crypto/bcrypt"

//Hash recebe uma string e retorna seu hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

//VerificarSenha compara uma senha e um hash, retornando se elas são iguais
func VerificarSenha(senhaHash, senhaString string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaString))
}
