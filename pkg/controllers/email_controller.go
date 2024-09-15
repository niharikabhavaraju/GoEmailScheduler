package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/niharikabhavaraju/go_scheduler/pkg/models"
	"github.com/niharikabhavaraju/go_scheduler/pkg/utils"
)

var NewEmail models.Email

func GetEmails(w http.ResponseWriter, r *http.Request) {
	newEmails, err := models.GetEmails()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	res, err := json.Marshal(newEmails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetEmailById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	emailId := vars["id"]
	ID, err := strconv.Atoi(emailId)
	if err != nil {
		fmt.Println("Error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	emailDetails, _, err := models.GetEmailById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	res, err := json.Marshal(emailDetails)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateEmail(w http.ResponseWriter, r *http.Request) {
	CreateEmail := &models.Email{}
	utils.ParseBody(r, CreateEmail)
	e, err := CreateEmail.CreateEmail()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	res, err := json.Marshal(e)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	emailId := vars["id"]
	ID, err := strconv.Atoi(emailId)
	if err != nil {
		fmt.Println("Error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	email, err := models.DeleteEmail(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	res, err := json.Marshal(email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	emailId := vars["id"]
	ID, err := strconv.Atoi(emailId)
	if err != nil {
		fmt.Println("Error while parsing")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}
	updatedEmail := &models.Email{}
	utils.ParseBody(r, updatedEmail)
	emailDetails, db, err := models.GetEmailById(ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	} else {
		if updatedEmail.Status != "" {
			emailDetails.Status = updatedEmail.Status
		}
		if updatedEmail.Subject != "" {
			emailDetails.Subject = updatedEmail.Subject
		}
		if updatedEmail.Body != "" {
			emailDetails.Body = updatedEmail.Body
		}
		if updatedEmail.To != "" {
			emailDetails.To = updatedEmail.To
		}
		if updatedEmail.Time.IsZero() {
			emailDetails.Time = updatedEmail.Time
		}
		saveRes := db.Save(&emailDetails)
		if saveRes.Error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		res, err := json.Marshal(emailDetails)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}
