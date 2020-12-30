package router

import (
	"net/http"

	"github.com/1gkx/taskmanager/internal/template"
)

func projectview(w http.ResponseWriter, r *http.Request) {
	template.Templates.ExecuteTemplate(w, "projectview", nil)
}

func projecttasks(w http.ResponseWriter, r *http.Request) {
	template.Templates.ExecuteTemplate(w, "projecttasks", nil)
}

func projectsetting(w http.ResponseWriter, r *http.Request) {
	template.Templates.ExecuteTemplate(w, "projectsettings", nil)
}
