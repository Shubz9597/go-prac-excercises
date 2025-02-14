package controller

import (
	"adventureGame/handler"
	"adventureGame/templates"
	"net/http"
)

func GetPageTemplate(w http.ResponseWriter, r *http.Request) {
	param := r.PathValue("title")

	story, exists := handler.ParsedData[param]

	if !exists {
		http.Error(w, "No Story found", http.StatusNotFound)
	}

	t, err := templates.CreateTemplate("template")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	t.Execute(w, story)

}
