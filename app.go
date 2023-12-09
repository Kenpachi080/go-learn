package main

import (
	"log"
	"net/http"
	"project-go/cmd/routes"
)

func main() {
	log.Println("Запуск веб-сервера на http://127.0.0.1:4000")

	err := http.ListenAndServe(":4000", routes.Init())
	log.Fatal(err)
}
