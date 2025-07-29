package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/http/handlers"
	"api/src/http/routers"
	"api/src/storage"
)

func main() {
	db, err := storage.InitDB("todos.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handler := &handlers.Handler{DB: db}

	router := routers.TodoRouter(handler)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
