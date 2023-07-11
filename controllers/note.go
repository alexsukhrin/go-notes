package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/alexsukhrin/go-notes/models"
	"github.com/alexsukhrin/go-notes/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func GetAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var notes []models.Note
	models.DB.Find(&notes)

	json.NewEncoder(w).Encode(notes)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var note models.Note

	if err := models.DB.Where("id = ?", id).First(&note).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	json.NewEncoder(w).Encode(note)
}

var validate *validator.Validate

type NoteInput struct {
	Title string `json:"title" validate:"required"`
	Body  string `json:"body" validate:"required"`
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var input NoteInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	note := &models.Note{
		Title: input.Title,
		Body:  input.Body,
	}

	models.DB.Create(note)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(note)

}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var note models.Note

	if err := models.DB.Where("id = ?", id).First(&note).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	var input NoteInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	note.Title = input.Title
	note.Body = input.Body

	models.DB.Save(&note)

	json.NewEncoder(w).Encode(note)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var note models.Note

	if err := models.DB.Where("id = ?", id).First(&note).Error; err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Note not found")
		return
	}

	models.DB.Delete(&note)

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(note)
}
