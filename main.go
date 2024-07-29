package main

import (
	"fmt"
	"log"
	"mongo-api/router"
	"net/http"
)

func main() {
	fmt.Println("Mongodb api")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Server is running at port 4000")
}
