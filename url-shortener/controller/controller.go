package controller

import (
	"fmt"
	"net/http"
	"urlShortner/handlers"
)

func GetRedirectUrl(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("code")

	config := handlers.Config
	fmt.Println(config)
	if param == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found"))
	}

	//We will find the key that is present
	for _, val := range config.Paths {
		if val.Path == param {
			http.Redirect(w, r, val.Url, http.StatusSeeOther)
		}
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Url not found"))
}
