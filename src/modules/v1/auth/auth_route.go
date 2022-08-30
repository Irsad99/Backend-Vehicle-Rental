package auth

import (
	"BackendGo/src/modules/v1/users"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth").Subrouter()

	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.Signin).Methods("POST")
}
