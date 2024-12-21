package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"memento/appContext"
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
		RespondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var memento = models.Memento{
		UserID: uint(userid),
	}

	mementos, err := memento.GetMementosByUserId(appContext.DB)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJson(w, http.StatusOK, mementos)
}

func CreateMemento(w http.ResponseWriter, r *http.Request) {
	var m models.Memento

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&m); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err.Error())
			RespondWithError(w, http.StatusInternalServerError, "Unexpected error occurred, try again")
		}
	}(r.Body)

	if err := m.CreateMemento(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithJson(w, http.StatusCreated, m)
}
