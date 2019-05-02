package movierouter

import (
	"encoding/json"
	"net/http"

	. "github.com/mguilhermetavares/poc-go-rest/config/dao"
	. "github.com/mguilhermetavares/poc-go-rest/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = BusinessDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	business, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, business)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	business, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Business ID")
		return
	}
	respondWithJson(w, http.StatusOK, business)
}

func Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var business Business
	if err := json.NewDecoder(r.Body).Decode(&business); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	business.ID = bson.NewObjectId()
	if err := dao.Create(business); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, business)
}

func Update(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	var business Business
	if err := json.NewDecoder(r.Body).Decode(&business); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(params["id"], business); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": business.Name + " atualizado com sucesso!"})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	params := mux.Vars(r)
	if err := dao.Delete(params["id"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
