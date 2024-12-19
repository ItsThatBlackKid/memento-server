package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"log"
	"memento/dto"
	"memento/models"
	"memento/utils"
	"net/http"
	"strconv"
)

type UserController struct {
	DB *sql.DB
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
	}

	u := models.User{ID: int8(int16(id))}
	if err := u.GetUser(uc.DB); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			respondWithError(w, http.StatusNotFound, "User not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, u.ToDTO())
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r.Body)

	if err := u.CreateUser(uc.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u.ToDTO())
}

func (uc *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
	}

	var u models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Something went wrong, please try again.")
		}
	}(r.Body)
	u.ID = int8(id)

	if err := u.UpdateUser(uc.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u.ToDTO())
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	type LoginResponse struct {
		Message string `json:"message"`
		Token   string `json:"token"`
	}

	// max 1MB size
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var u dto.LoginUser
	var user models.User
	if err := dec.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User payload")
		return
	}

	log.Println("decoded user: ", u)

	if err := user.LoginUser(uc.DB, u); err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}

	token, err := utils.EncodeJwt(user.ToDTO())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong, please try again.")
		return
	}

	resp := LoginResponse{
		Message: "Login successful!",
		Token:   token,
	}

	respondWithJSON(w, http.StatusAccepted, resp)
}
