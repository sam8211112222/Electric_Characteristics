package main

import (
	"Electric_Characteristics/config"
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	var err error

	fd := returnExcelData(config.FilePath, config.SheetIndex)

	mongo, ctx, err := connectToMongo(config.Uri, &config.Cred, config.MongoClient)

	defer CloseClientDB(mongo, ctx)

	dataCollection := returnCollection(mongo, config.SpiceDB, config.DataTable)

	// truncate data
	dropCollection(dataCollection, ctx)

	// insert data
	res := insertOne(dataCollection, ctx, fd)
	id := res.InsertedID

	// select data
	t := Fulldata{}
	if err = findOne(dataCollection, ctx, id).Decode(&t); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", t))

}

func (fd Fulldata) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(fd)
	if err != nil {
		return
	}
}
