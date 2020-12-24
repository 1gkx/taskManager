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

	route.HandleFunc("/", index).Methods("GET")
	// r.Handle("/user", authRequireHandlerWrap(profile)).Methods("GET")

	// r.NotFoundHandler = http.HandlerFunc(errorhendler)

	return route
}

// func setPublicFolder() {
// 	publicFolder := http.FileServer(http.Dir("./public"))
// 	r.PathPrefix("/js/").Handler(publicFolder)
// 	r.PathPrefix("/css/").Handler(publicFolder)
// 	r.PathPrefix("/img/").Handler(publicFolder)
// }

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
