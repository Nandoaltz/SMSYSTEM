package hashsenha

import "golang.org/x/crypto/bcrypt"

func ConvertHash(senha string)([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func CompareHash(senhaNormal string, senhaHash string)error{
	return bcrypt.CompareHashAndPassword([]byte(senhaHash), []byte(senhaNormal))
}