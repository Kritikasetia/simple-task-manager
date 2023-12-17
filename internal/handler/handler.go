package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"simpletaskmanager/internal/app"
	"strconv"
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

func ReadTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := r.URL.Query().Get("id")
	if taskIDStr == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	task, err := app.ReadTask(context.Background(), db, taskID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading task: %v", err), http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	err = app.UpdateTask(context.Background(), db, task)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error updating task: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Task updated successfully")
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	taskIDStr := r.URL.Query().Get("id")
	if taskIDStr == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid Task ID", http.StatusBadRequest)
		return
	}

	err = app.DeleteTask(context.Background(), db, taskID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting task: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Task deleted successfully")
}
