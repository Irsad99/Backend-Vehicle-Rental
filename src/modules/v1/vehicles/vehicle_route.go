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

	route.HandleFunc("/", middleware.Do(ctrl.GetAll, "admin", middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/product", middleware.Do(ctrl.GetByID, "user", middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/search", middleware.Do(ctrl.SearchByType, "user", middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/sort", middleware.Do(ctrl.SortByPLT, "user", middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/popular", ctrl.PopularVehicle).Methods("GET")
	route.HandleFunc("/add", middleware.Do(ctrl.Delete, "admin", middleware.CheckAuth)).Methods("POST")
	route.HandleFunc("/delete/{id}", middleware.Do(ctrl.Delete, "admin", middleware.CheckAuth)).Methods("DELETE")
	route.HandleFunc("/update", middleware.Do(ctrl.Update, "admin", middleware.CheckAuth)).Methods("PUT")
}
