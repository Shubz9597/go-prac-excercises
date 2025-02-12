package routes

import (
	"net/http"
	"urlShortner/controller"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/{code}", controller.GetRedirectUrl)

	return router
}
