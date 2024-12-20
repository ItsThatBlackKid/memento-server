package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"memento/context"
	"memento/models"
	"net/http"
	"strconv"
)

type MementoController struct {
	DB *sql.DB
}

func GetMementos(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid, err := strconv.Atoi(vars["userid"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var memento = models.Memento{
		UserID: userid,
	}

	mementos, err := memento.GetMementosByUserId(context.Context.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, mementos)
}

func CreateMemento(w http.ResponseWriter, r *http.Request) {
	var m models.Memento

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&m); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	if err := m.CreateMemento(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, m)
}
