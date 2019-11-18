package search

import (
	"encoding/json"
	"net/http"

	"booking-service/commons"

	"github.com/gorilla/mux"
)

var RoomDetail = []*commons.Room{}
var FreeRoomList = []*commons.Room{}

func Handler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var roomList = []*commons.Room{}
	if id == "" {
		roomList = searchAll()
	} else {
		roomList = searchByID(id)
	}

	w.WriteHeader(http.StatusOK)
	res := &commons.Response{}
	res.Success(id+" is OK", roomList)
	json.NewEncoder(w).Encode(res)
}

func searchAll() []*commons.Room {
	//read and calculate free room
	return nil
}

func searchByID(id string) []*commons.Room {

	return nil
}
