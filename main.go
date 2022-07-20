package main

import (
	"Electric_Characteristics/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	defer CloseClientDB(mongoConnection, ctx)

	// truncate data if need
	dropCollection(mongoConnection, config.SpiceDB, config.DataTable, ctx)

	// insert data
	_, err := insertData(mongoConnection, config.SpiceDB, config.DataTable, ctx, fd)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = config.DefaultPort
	}

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/search/{id}", getOneData).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+config.DefaultPort, router))

}

func getOneData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	queryId := vars["id"] //獲取url參數
	result, err := findDataById(mongoConnection, config.SpiceDB, config.DataTable, queryId)
	if err != nil {
		log.Fatal(err)
	}
	t := &Fulldata{}
	err = result.Decode(t)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(t)
	if err != nil {
		return
	}
}
