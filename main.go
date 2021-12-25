package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	buildDir := http.Dir("./client/build")

	http.Handle("/", http.FileServer(buildDir))
	log.Println("Listening on PORT :", port)
	httpErr := http.ListenAndServe(":"+port, nil)

	if httpErr != nil {
		panic("Error: " + httpErr.Error())
	}

}
