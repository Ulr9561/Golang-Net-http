package handlers

import (
	"My_first_go_project/internal/models"
	"My_first_go_project/internal/storage"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	db *storage.Database
}

func NewTaskHandler(db *storage.Database) *TaskHandler {
	return &TaskHandler{db: db}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.db.GetTasks()
	if err != nil {
		http.Error(w, "Erreur lors de la récuperation des taches", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(tasks)
	if err != nil {
		return
	}
}

func (h *TaskHandler) AddTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	if err := h.db.AddTask(&task); err != nil {
		http.Error(w, "Erreur lors de l'ajout", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		return
	}
}
