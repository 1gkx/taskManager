package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/1gkx/taskmanager/internal/store"
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

func projectadd(w http.ResponseWriter, r *http.Request) {

	pr := new(store.Project)
	if err := json.NewDecoder(r.Body).Decode(&pr); err != nil {
		RespAPI(http.StatusInternalServerError, w, err)
		return
	}

	fmt.Printf("Project: %+v\n", pr)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
		"data":   pr,
	})
}
