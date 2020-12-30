package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/1gkx/taskmanager/internal/store"
	"github.com/1gkx/taskmanager/internal/template"
	"github.com/gorilla/mux"
)

func userview(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	u := store.FindByID(uint(id))

	isNew := false
	if uint(id) == 0 {
		isNew = true
	}

	w.Header().Set("Cache-Control", "No-Cache")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Templates.ExecuteTemplate(w, "user",
		map[string]interface{}{
			"title":     fmt.Sprintf("Admin / %s", u.DisplayName()),
			"isNew":     isNew,
			"isProfile": false,
			"user":      nil,
			"data":      u,
			"redirect":  "/admin/users",
		},
	)
	return

}

func userList(w http.ResponseWriter, r *http.Request) {

	page, limit := 1, 10

	params := r.URL.Query()
	if len(params["page"]) > 0 {
		page, _ = strconv.Atoi(params["page"][0])
	}
	if len(params["limit"]) > 0 {
		limit, _ = strconv.Atoi(params["limit"][0])
	}

	template.Templates.ExecuteTemplate(w, "list", map[string]interface{}{
		"title": "Admin / Users",
		"user":  nil,
		"data":  store.FindUser(page, limit),
	})
}

func userAdd(w http.ResponseWriter, r *http.Request) {

	// bodyBytes, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	RespAPI(http.StatusInternalServerError, w, err.Error())
	// 	return
	// }
	// bodyString := string(bodyBytes)
	// fmt.Printf("Body: %s\n", bodyString)

	u := new(store.User)
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		RespAPI(http.StatusInternalServerError, w, err)
		return
	}

	fmt.Printf("User: %+v\n", u)

	// New User
	if u.ID == 0 {
		if err := store.CreateOrUpdate(u); err != nil {
			RespAPI(http.StatusInternalServerError, w, err.Error())
			return
		}
		return
	}
	if u.ID > 0 {
		if err := store.CreateOrUpdate(u); err != nil {
			RespAPI(http.StatusInternalServerError, w, err.Error())
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": false,
		"data":   "Invalid or missing id",
	})
	return
}

func RespAPI(code int, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": false,
		"data":   data,
	})
}
