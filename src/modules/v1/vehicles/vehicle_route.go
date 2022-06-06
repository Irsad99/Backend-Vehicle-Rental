package vehicles

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicle").Subrouter()

	repo := NewRepo(db)
	ctrl := NewCtrl(repo)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/search", ctrl.SearchByType).Methods("GET")
	route.HandleFunc("/price", ctrl.SortByPrice).Methods("GET")
	route.HandleFunc("/popular", ctrl.PopularVehicle).Methods("GET")
	route.HandleFunc("/register", ctrl.AddData).Methods("POST")
	route.HandleFunc("/delete/{id}", ctrl.Delete).Methods("DELETE")
	route.HandleFunc("/update", ctrl.Update).Methods("PUT")
}