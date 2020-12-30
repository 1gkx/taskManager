package router

import (
	"net/http"

	"github.com/1gkx/taskmanager/internal/template"
)

func dashboard(w http.ResponseWriter, r *http.Request) {
	template.Templates.ExecuteTemplate(w, "admin.dashboard", nil)
}
