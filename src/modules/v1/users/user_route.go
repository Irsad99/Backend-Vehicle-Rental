package users

import (
	"BackendGo/src/middleware"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/", middleware.Do(ctrl.GetAll, "admin", middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/register", ctrl.AddData).Methods("POST")
	route.HandleFunc("/delete/{id}", middleware.Do(ctrl.Delete, "admin", middleware.CheckAuth)).Methods("DELETE")
	route.HandleFunc("/update", middleware.Do(ctrl.Update, "user", middleware.CheckAuth)).Methods("PUT")
}
