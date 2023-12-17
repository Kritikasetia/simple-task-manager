package main

import (
	"fmt"
	"net/http"
	"simpletaskmanager/internal/db"
	"simpletaskmanager/internal/handler"
)

func main() {
	db.InitDB()
	defer db.CloseDB()

	handler.InitHandler(db.GetDB())

	http.HandleFunc("/tasks/create", handler.CreateTaskHandler)
	http.HandleFunc("/tasks/read", handler.ReadTaskHandler)
	http.HandleFunc("/tasks/update", handler.UpdateTaskHandler)
	http.HandleFunc("/tasks/delete", handler.DeleteTaskHandler)

	port := 8080
	serverAddr := fmt.Sprintf(":%d", port)

	fmt.Printf("Server listening on %s...\n", serverAddr)
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		panic(err)
	}

}
