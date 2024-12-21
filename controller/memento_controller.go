package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"memento/appContext"
	"memento/models"
	"net/http"
)

type MementoController struct {
	DB *sql.DB
}
type UserIDHandlerFunc func(userID uint)

func HandleWithUserID(w http.ResponseWriter, r *http.Request, handler UserIDHandlerFunc) {
	ctx := r.Context()
	userID, ok := ctx.Value("user_id").(uint)

	if !ok {
		RespondWithError(w, http.StatusUnauthorized, "Unauthorized Access: invalid or missing bearer token")
		return
	}

	handler(userID)
}

func GetMementos(w http.ResponseWriter, r *http.Request) {
	HandleWithUserID(
		w, r, func(userID uint) {
			var memento = models.Memento{
				UserID: userID,
			}

			mementos, err := memento.GetMementosByUserId(appContext.DB)
			if err != nil {
				RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}

			ResponseWithJson(w, http.StatusOK, mementos)
		},
	)
}

func CreateMemento(w http.ResponseWriter, r *http.Request) {
	HandleWithUserID(
		w, r, func(userID uint) {
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

			m.UserID = userID

			if err := m.CreateMemento(); err != nil {
				RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}

			ResponseWithJson(w, http.StatusCreated, m)
		},
	)
}
