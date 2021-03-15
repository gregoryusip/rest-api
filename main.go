package main

import (
	"fmt"
	"log"
	"net/http"
	"rest-api/router"
)

func main() {
	r := router.Router()
	fmt.Println("Server dijalankan pada port 7070...")

	log.Fatal(http.ListenAndServe(":7070", r))
}
