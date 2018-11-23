package routes

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"encoding/json"
	"dellemc/system-state-service/src/db"
)


func GetAssets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	Assets := db.GetAssets()
	fmt.Println("Assets Loaded")
	json.NewEncoder(w).Encode(Assets)
	

}

func GetSystemState( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "application/json")
 	systemstate := db.GetSystemState()
	fmt.Println("State Loaded")
	json.NewEncoder(w).Encode(systemstate)
}

//Comment
func NewRouter() *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/assets", GetAssets).Methods("GET")	
	router.HandleFunc("/api/systemstate", GetSystemState).Methods("GET")	
	return router
}

