package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/1gkx/taskmanager/internal/template"
)

var route *mux.Router

func NewRouter() *mux.Router {

	route = mux.NewRouter().StrictSlash(true)

	// setPublicFolder()
	// setUserRouters()
	// setSignInRouters()
	// setSettingRouters()

	publicFolder := http.FileServer(http.Dir("./public"))
	route.PathPrefix("/js/").Handler(publicFolder)
	route.PathPrefix("/css/").Handler(publicFolder)
	route.PathPrefix("/img/").Handler(publicFolder)

	// route.NotFoundHandler = http.HandlerFunc(errorhendler)
	route.HandleFunc("/", index).Methods("GET")
	// r.Handle("/user", authRequireHandlerWrap(profile)).Methods("GET")

	// Users
	// route.Handle("/admin/users", authRequireHandlerWrap(userList)).Methods("GET")
	route.HandleFunc("/admin/users", userList).Methods("GET")
	route.HandleFunc("/admin/users/{id:[0-9]+}", userprofile).Methods("GET")
	route.HandleFunc("/admin/users", userAdd).Methods("POST")

	// 	r.Handle("/admin/users/{id:[0-9]+}", authRequireHandlerWrap(userprofile)).Methods("GET")
	// 	r.Handle("/admin/users", authRequireHandlerWrap(userAdd)).Methods("POST")
	// 	r.Handle("/admin/users", authRequireHandlerWrap(userUpdate)).Methods("PUT")
	// 	r.Handle("/admin/users", authRequireHandlerWrap(userRemove)).Methods("DELETE")

	return route
}

// func index(next http.Handler) http.Handler {
//     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//         // Do stuff here
//         log.Println(r.RequestURI)
//         // Call the next handler, which can be another middleware in the chain, or the final handler.
//         next.ServeHTTP(w, r)
//     })
// }

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "No-Cache")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	template.Templates.ExecuteTemplate(w, "home", nil)
}

// func errorhendler(w http.ResponseWriter, r *http.Request) {
// 	responce(w, r, "status/40x", nil)
// }

// func profile(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	w.Header().Set("Cache-Control", "No-Cache")
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	templates.Templates.ExecuteTemplate(w, "user",
// 		map[string]interface{}{
// 			"isNew":     false,
// 			"isProfile": true,
// 			"user":      session.GetUser(r),
// 			"data":      session.GetUser(r),
// 			"redirect":  "/user",
// 		},
// 	)
// 	return
// }

// func responceJson(code int, w http.ResponseWriter, data map[string]interface{}) {
// 	w.WriteHeader(code)
// 	w.Header().Set("Cache-Control", "No-Cache")
// 	json.NewEncoder(w).Encode(map[string]interface{}{
// 		"code": code,
// 		"data": data,
// 	})
// }

// func responce(w http.ResponseWriter, r *http.Request, tmpl string, data interface{}) {
// 	w.Header().Set("Cache-Control", "No-Cache")
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	templates.Templates.ExecuteTemplate(w, tmpl,
// 		map[string]interface{}{
// 			"user": session.GetUser(r),
// 			"data": data,
// 		},
// 	)
// }
