package aut

import (
	"TCC/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	// "crypto/rand"
	// "encoding/base64"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// func GenerateSecretKey() (string, error) {
// 	key := make([]byte, 64)
// 	_, err := rand.Read(key)
// 	if err != nil {
// 		return "", err
// 	}
// 	secret := base64.URLEncoding.EncodeToString(key)
// 	return secret, nil
// }

func JsonWebToken(usuarioID uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    usuarioID,                    // ID do usuário
		"exp":   time.Now().Add(time.Hour * 5).Unix(), // Expiração do token
		"authorized": true,            
	})
	return token.SignedString(config.Signiture)
}

func ValidToken(r *http.Request) error {
	tokenString := extrairToken(r)

	//Essa é a função que verifica se o token não foi alterado
	tok3n, erro := jwt.Parse(tokenString, keyVerificacao)
	if erro != nil {
		return erro
	}
	//Uma verificação para ver se o conteudo das claims podem ser convertidas para um jwt.MapClaims
	//Que converte os dados salvos em uma interface genérica para uma estrutura de dados mais organizada que é o MapClaims 
	if _, status := tok3n.Claims.(jwt.MapClaims); status && tok3n.Valid {
		return nil
	}
	return errors.New("Token inválido")
}

func ExtrairIDDoToken(r *http.Request) (uint64, error) {
    token := extrairToken(r)
    if token == "" {
        return 0, errors.New("Token não encontrado na requisição")
    }
    tok3n, erro := jwt.Parse(token, keyVerificacao)
    if erro != nil {
        return 0, errors.New("Token invalido")
	}
    if claims, ok := tok3n.Claims.(jwt.MapClaims); ok && tok3n.Valid {
        if sub, ok := claims["sub"].(float64); ok {
            UserID := uint64(sub)
            if erro != nil {
                return 0, erro
            }
            return UserID, nil
        }
        return 0, errors.New("ID do token inválido")
    }

    return 0, errors.New("Erro na validação do token")
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func keyVerificacao(token *jwt.Token) (interface{}, error) {
	if _, status := token.Method.(*jwt.SigningMethodHMAC); !status {
		return nil, fmt.Errorf("%v", token.Header["alg"])
	}
	return config.Signiture, nil
}

