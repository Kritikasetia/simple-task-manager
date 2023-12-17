package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"simpletaskmanager/internal/app"
)

var db *sql.DB

func InitHandler(database *sql.DB) {
	db = database
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task app.Task

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	id, err := app.CreateTask(context.Background(), db, task)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Task created with ID: %d", id)
}
