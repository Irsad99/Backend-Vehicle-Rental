package users

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()

	repo := NewRepo(db)
	ctrl := NewCtrl(repo)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/register", ctrl.AddData).Methods("POST")
	route.HandleFunc("/delete/{id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/update", ctrl.Update).Methods("PUT")
}