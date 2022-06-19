package routers

import (
	"net/http"

	"BackendGo/src/database"
	"BackendGo/src/modules/v1/auth"
	"BackendGo/src/modules/v1/histories"
	"BackendGo/src/modules/v1/users"
	"BackendGo/src/modules/v1/vehicles"

	// "github.com/gorilla/handlers"
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
	auth.New(mainRoute, db)

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080/")
    // w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Write([]byte("{\"hello\": \"world\"}"))
}
