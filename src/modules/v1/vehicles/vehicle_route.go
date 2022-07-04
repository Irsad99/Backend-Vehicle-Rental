package vehicles

import (
	"BackendGo/src/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicle").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", ctrl.GetAll).Methods("GET")
	route.HandleFunc("/product", ctrl.GetByID).Methods("GET")
	route.HandleFunc("/search", ctrl.SearchByType).Methods("GET")
	route.HandleFunc("/sort", ctrl.SortByPLT).Methods("GET")
	route.HandleFunc("/popular", ctrl.PopularVehicle).Methods("GET")
	route.HandleFunc("/register", ctrl.AddData).Methods("POST")
	route.HandleFunc("/delete/{id}", middleware.Do(ctrl.Delete, "admin", middleware.CheckAuth)).Methods("DELETE")
	route.HandleFunc("/update", middleware.Do(ctrl.Update, "user", middleware.CheckAuth)).Methods("PUT")
}
