package test

import (
	"context"
	"simpletaskmanager/internal/app"
	"simpletaskmanager/internal/db"
	"testing"
)

func TestCreateTask(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	ctx := context.Background()

	// Ensure the tasks table is empty before testing
	_, err := db.GetDB().ExecContext(ctx, "DELETE FROM tasks")
	if err != nil {
		t.Fatal(err)
	}

	task := app.Task{
		Title:       "Test Task",
		Description: "This is a test task.",
	}

	id, err := app.CreateTask(ctx, db.GetDB(), task)
	if err != nil {
		t.Fatal(err)
	}

	if id == 0 {
		t.Fatal("Task ID should not be 0")
	}
}

func TestReadTask(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	ctx := context.Background()

	_, err := db.GetDB().ExecContext(ctx, "DELETE FROM tasks")
	if err != nil {
		t.Fatal(err)
	}
	task := app.Task{
		Title:       "Test Task",
		Description: "This is a test task.",
	}

	id, err := app.CreateTask(ctx, db.GetDB(), task)
	if err != nil {
		t.Fatal(err)
	}

	taskID := id // Assuming the ID of the task created above

	readTaskResult, err := app.ReadTask(ctx, db.GetDB(), taskID)
	if err != nil {
		t.Fatal(err)
	}

	if readTaskResult.ID != taskID {
		t.Fatalf("Expected task ID %d, got %d", taskID, task.ID)
	}
}

func TestUpdateTask(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	ctx := context.Background()

	// Ensure the tasks table is not empty before testing
	_, err := db.GetDB().ExecContext(ctx, "DELETE FROM tasks")
	if err != nil {
		t.Fatal(err)
	}
	task := app.Task{
		Title:       "Test Task",
		Description: "This is a test task.",
	}

	id, err := app.CreateTask(ctx, db.GetDB(), task)
	if err != nil {
		t.Fatal(err)
	}
	taskID := id // Assuming the ID of the task created above

	updatedTask := app.Task{
		ID:          taskID,
		Title:       "Updated Test Task",
		Description: "This is the updated test task.",
	}

	err = app.UpdateTask(ctx, db.GetDB(), updatedTask)
	if err != nil {
		t.Fatal(err)
	}

	// Retrieve the task after update
	retrievedTask, err := app.ReadTask(ctx, db.GetDB(), taskID)
	if err != nil {
		t.Fatal(err)
	}

	if retrievedTask.Title != updatedTask.Title || retrievedTask.Description != updatedTask.Description {
		t.Fatalf("Update failed. Expected: %+v, Got: %+v", updatedTask, retrievedTask)
	}
}

func TestDeleteTask(t *testing.T) {
	db.InitDB()
	defer db.CloseDB()

	ctx := context.Background()

	// Ensure the tasks table is not empty before testing
	_, err := db.GetDB().ExecContext(ctx, "DELETE FROM tasks")
	if err != nil {
		t.Fatal(err)
	}
	task := app.Task{
		Title:       "Test Task",
		Description: "This is a test task.",
	}

	id, err := app.CreateTask(ctx, db.GetDB(), task)
	if err != nil {
		t.Fatal(err)
	}

	taskID := id // Assuming the ID of the task created above

	err = app.DeleteTask(ctx, db.GetDB(), taskID)
	if err != nil {
		t.Fatal(err)
	}

	// Attempt to retrieve the task after delete
	_, err = app.ReadTask(ctx, db.GetDB(), taskID)
	if err == nil {
		t.Fatal("Expected error when reading a deleted task, but got none.")
	}
}
