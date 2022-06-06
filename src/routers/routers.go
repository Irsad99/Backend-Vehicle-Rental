package routers

import (
	"net/http"

	"BackendGo/src/configs/database"
	"BackendGo/src/modules/v1/users"
	"BackendGo/src/modules/v1/vehicles"
	"BackendGo/src/modules/v1/histories"

	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	histories.New(mainRoute, db)

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello worlds"))
}
