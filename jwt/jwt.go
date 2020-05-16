package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gonza213/twitter/models"
)

//GeneroJWT
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte("Bobby2020")

	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellido":        t.Apellido,
		"fechaNacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"SitioWeb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
