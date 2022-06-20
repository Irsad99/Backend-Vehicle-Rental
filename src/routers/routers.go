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
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	nRelic, err := newrelic.NewApplication(
		newrelic.ConfigAppName("BackendGo"),
		newrelic.ConfigLicense("7587a4727585c6b87bf5a5935dcf8c76a5eeNRAL"),
		newrelic.ConfigDistributedTracerEnabled(true),
	)

	mainRoute.Use(nrgorilla.Middleware(nRelic))

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	mainRoute.HandleFunc(newrelic.WrapHandleFunc(nRelic, "/", relicHandler)).Methods("GET")
	users.New(mainRoute, db)
	vehicles.New(mainRoute, db)
	histories.New(mainRoute, db)
	auth.New(mainRoute, db)

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"hello\": \"world\"}"))
}

func relicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("{\"hello\": \"world\"}"))
}
