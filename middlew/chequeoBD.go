package middlew

import (
	"net/http"

	"github.com/gonza213/twitter/bd"
)

// ChequeoBD
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión pérdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
