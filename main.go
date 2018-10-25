package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const message = `Hi gogoconf 2018, the first Golang conference in Poland! It's 3:10pm 😬`

func mainHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, message)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainHandler)

	port := ":8081"
	if v, ok := os.LookupEnv("PORT"); ok {
		port = ":" + v
	}
	log.Println("Listening on " + port)
	log.Fatal(http.ListenAndServe(port, r))
}
