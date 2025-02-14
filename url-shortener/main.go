package main

import (
	"flag"
	"fmt"
	"net/http"
	"urlShortner/handlers"
	"urlShortner/routes"

	_ "github.com/lib/pq"
)

type ProgramParams struct {
	yaml bool
	file string
}

func main() {
	var params ProgramParams
	flag.BoolVar(&params.yaml, "yaml", false, "Takes yaml file")
	flag.StringVar(&params.file, "file", "path", "The file that is needed to be parsed")

	flag.Parse()

	handlers.ParseFileAndReturn(params.file, params.yaml)

	mux := routes.Router()

	fmt.Println("Server started at port 3001")
	http.ListenAndServe(":3001", mux)

}
