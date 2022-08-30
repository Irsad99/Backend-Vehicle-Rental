package auth

import (
	"BackendGo/src/database/gorm/models"
	"BackendGo/src/interfaces"
	"encoding/json"
	"net/http"
	
)
type Auth_ctrl struct {
	svc interfaces.AuthService
}

func NewCtrl(ctrl interfaces.AuthService) *Auth_ctrl {
	return &Auth_ctrl{ctrl}
}

func (ctrl *Auth_ctrl) Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.User

	json.NewDecoder(r.Body).Decode(&data)

	result, err := ctrl.svc.Login(data)
	if err != nil {
		response.ResponseJSON(400, "Login Gagal").Send(w)
		return
	}

	json.NewEncoder(w).Encode(&result)

}
