package main

import (
	"adventureGame/handler"
	"adventureGame/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("adeventrure game")

	jsonData, err := handler.ParseJson()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(jsonData["debate"].Title)

	//Now we create the mux

	mux := routes.Router()
	fmt.Print("Server Starting in 3001")
	http.ListenAndServe(":3001", mux)
}
