package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handler("/add-menu", http.HandlerFunc(Handler.AddMenu))

	fmt.Println("menu service listen on port : 8000")
	log.Panic(http.ListenAndServe(":8000", router))
}
