package router

import (
	"net/http"

	"github.com/1gkx/taskmanager/internal/template"
)

func hometask(w http.ResponseWriter, r *http.Request) {
	template.Templates.ExecuteTemplate(w, "home.tasks", nil)
}

func homeproject(w http.ResponseWriter, r *http.Request) {
	template.Templates.ExecuteTemplate(w, "home.projects", nil)
}

func hometimeline(w http.ResponseWriter, r *http.Request) {
	template.Templates.ExecuteTemplate(w, "home.timeline", nil)
}
