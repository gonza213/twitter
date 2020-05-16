package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gonza213/twitter/bd"
	"github.com/gonza213/twitter/models"
)

//Registro
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 500)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "La contraseña debe tener minimo 6 digitos", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado == true {
		http.Error(w, "Ya existe usuario con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al ingresar el registro de usuarios"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrador registrar el ususario"+err.Error(), 400)
	}

	w.WriteHeader(http.StatusCreated)
}
