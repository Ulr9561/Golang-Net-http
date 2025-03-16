package main

import (
	"My_first_go_project/internal/handlers"
	"My_first_go_project/internal/storage"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db := storage.Init("tasks.db")

	taskHandler := handlers.NewTaskHandler(db)

	r := mux.NewRouter()

	r.HandleFunc("/api/tasks", taskHandler.GetTasks).Methods("GET")
	r.HandleFunc("/api/tasks", taskHandler.AddTask).Methods("POST")

	fmt.Println("🚀 Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
