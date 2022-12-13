package main

import (
	"fmt"
	"log"
	"net/http"
	"todoapp/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting the server on port 9000...")

	log.Fatal(http.ListenAndServe(":9000", r))
}