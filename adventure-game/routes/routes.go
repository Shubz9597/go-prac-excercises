package routes

import (
	"adventureGame/controller"
	"adventureGame/templates"
	"net/http"
)

func Router() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/story/", func(w http.ResponseWriter, r *http.Request) {
		t, err := templates.CreateTemplate("start")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		t.Execute(w, "")
	})
	router.HandleFunc("/story/{title}", controller.GetPageTemplate)

	return router
}
