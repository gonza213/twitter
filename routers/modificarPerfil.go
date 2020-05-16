package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gonza213/twitter/bd"
	"github.com/gonza213/twitter/models"
)

//ModificarRegistro
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
