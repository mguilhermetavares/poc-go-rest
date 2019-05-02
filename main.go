package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/mguilhermetavares/poc-go-rest/config"
	. "github.com/mguilhermetavares/poc-go-rest/config/dao"
	businessrouter "github.com/mguilhermetavares/poc-go-rest/router"
)

var dao = BusinessDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/business", businessrouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/business/{id}", businessrouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/business", businessrouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/business/{id}", businessrouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/business/{id}", businessrouter.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
