package vehicles

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"BackendGo/src/database/gorm/models"
	"BackendGo/src/interfaces"

	"github.com/gorilla/mux"
)

type vehicle_ctrl struct {
	svc interfaces.VehicleService
}

func NewCtrl(ctrl interfaces.VehicleService) *vehicle_ctrl {
	return &vehicle_ctrl{ctrl}
}

func (ctrl *vehicle_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.FindAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *vehicle_ctrl) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataID = r.URL.Query()
	id, err := strconv.Atoi(dataID["id"][0])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	data, err := ctrl.svc.FindByID(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *vehicle_ctrl) SearchByType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = r.URL.Query()
	dataType := string(data["category"][0])
	dataLocation := string(data["location"][0])

	result, err := ctrl.svc.Search(&dataType, &dataLocation)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(result)
}

func (ctrl *vehicle_ctrl) SortByPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataPrice = r.URL.Query()
	price, err := strconv.Atoi(dataPrice["price"][0])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	data, err := ctrl.svc.SortByPrice(price)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *vehicle_ctrl) PopularVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataRating = r.URL.Query()
	rating, err := strconv.Atoi(dataRating["rating"][0])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	data, err := ctrl.svc.Popular(rating)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *vehicle_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.Vehicle
	json.NewDecoder(r.Body).Decode(&data)

	result, err := ctrl.svc.Save(&data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (ctrl *vehicle_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	id, err := strconv.Atoi(data["id"])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	result, err := ctrl.svc.Delete(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (ctrl *vehicle_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataId = r.URL.Query()
	var data models.Vehicle
	var reqId = r.Header.Get("id")
	var reqRole = r.Header.Get("role")

	json.NewDecoder(r.Body).Decode(&data)

	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	result, err := ctrl.svc.Update(id, &data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	if reqId != dataId["id"][0] {
		if reqRole == "admin" {
			return
		} else {
			response.ResponseJSON(401, "Akses Tidak Diijinkan").Send(w)
			return
		}
	}

	json.NewEncoder(w).Encode(&result)
}
