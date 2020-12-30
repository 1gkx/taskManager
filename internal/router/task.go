package router

import (
	"net/http"

	"github.com/1gkx/taskmanager/internal/template"
)

func taskview(w http.ResponseWriter, r *http.Request) {
	template.Templates.ExecuteTemplate(w, "task.view", nil)
}
