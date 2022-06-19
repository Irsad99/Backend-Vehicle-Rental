package histories

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"BackendGo/src/database/gorm/models"
	"BackendGo/src/interfaces"

	"github.com/gorilla/mux"
)

type history_ctrl struct {
	svc interfaces.HistoryService
}

func NewCtrl(ctrl interfaces.HistoryService) *history_ctrl {
	return &history_ctrl{ctrl}
}

func (ctrl *history_ctrl) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.FindAll()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *history_ctrl) SortByStartDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data, err := ctrl.svc.SortByStart()
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(data)
}

func (ctrl *history_ctrl) GetHistoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	var reqId = r.Header.Get("id")
	var reqRole = r.Header.Get("role")

	id, err := strconv.Atoi(data["id"])
	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	result, err := ctrl.svc.FindHistoryByID(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	if reqId != string(data["id"][0]){
		if reqRole == "admin" {
			json.NewEncoder(w).Encode(result)
			return
		} else {
			response.ResponseJSON(401, "Akses Tidak Diijinkan").Send(w)
			return
		}
	}

	json.NewEncoder(w).Encode(result)
}

func (ctrl *history_ctrl) AddData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data models.History
	json.NewDecoder(r.Body).Decode(&data)

	result, err := ctrl.svc.Save(&data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (ctrl *history_ctrl) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data = mux.Vars(r)
	id, err := strconv.Atoi(data["id"])

	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	result, err := ctrl.svc.Delete(id)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}

func (ctrl *history_ctrl) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var dataId = r.URL.Query()
	var status = string(dataId["status"][0])
	var reqId = r.Header.Get("id")
	var reqRole = r.Header.Get("role")
	
	if reqId != dataId["id"][0]{
		if reqRole == "admin" {
			return
		} else {
			response.ResponseJSON(401, "Akses Tidak Diijinkan").Send(w)
			return
		}
	}

	id, err := strconv.Atoi(dataId["id"][0])
	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	}

	result, err := ctrl.svc.Update(id, status)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	json.NewEncoder(w).Encode(&result)
}
