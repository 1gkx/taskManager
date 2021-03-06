package template

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/1gkx/taskmanager/internal/store"
)

var Templates *template.Template

// var funcMap = template.FuncMap{
// 	"users": store.UserCount(),
// 	"pagginated": func(arr interface{}) bool {
// 		arrType := reflect.TypeOf(arr).String()
// 		if arrType == "[]*store.Client" {
// 			return 5 > 2
// 		}
// 		if arrType == "[]*store.User" {
// 			return len(store.FindUser(1, 10)) > 2
// 		}
// 		return false
// 	},
// 	"listPage": func(arr interface{}) []struct{} {
// 		var arrLength int
// 		arrType := reflect.TypeOf(arr).String()
// 		if arrType == "[]*store.Client" {
// 			arrLength = 5
// 		}
// 		if arrType == "[]*store.User" {
// 			arrLength = len(store.FindUser(1, 10))
// 		}
// 		// Вынести смещение в переменную, т.к. используется в нескольких местах
// 		count := arrLength / 5
// 		if arrLength-count*5 > 0 {
// 			count = count + 1
// 		}
// 		return make([]struct{}, count)
// 	},
// 	"inc": func(i int) int {
// 		return i + 1
// 	},
// 	"copyrightYear": func() string {
// 		return fmt.Sprintf("%d", time.Now().Year())
// 	},
// }

func InitTemplate() {

	var pathTemplates = "templates/*"
	var templates []string

	tmpls := getFiles(pathTemplates, &templates)

	Templates = template.Must(template.New("").Funcs(
		template.FuncMap{
			"users": func() int64 {
				return store.UserCount()
			},
			"pagginated": func(arr interface{}) bool {
				arrType := reflect.TypeOf(arr).String()
				if arrType == "[]*store.Client" {
					return 5 > 2
				}
				if arrType == "[]*store.User" {
					return len(store.FindUser(1, 10)) > 2
				}
				return false
			},
			"listPage": func(arr interface{}) []struct{} {
				var arrLength int
				arrType := reflect.TypeOf(arr).String()
				if arrType == "[]*store.Client" {
					arrLength = 5
				}
				if arrType == "[]*store.User" {
					arrLength = len(store.FindUser(1, 10))
				}
				// Вынести смещение в переменную, т.к. используется в нескольких местах
				count := arrLength / 5
				if arrLength-count*5 > 0 {
					count = count + 1
				}
				return make([]struct{}, count)
			},
			"inc": func(i int) int {
				return i + 1
			},
			"copyrightYear": func() string {
				return fmt.Sprintf("%d", time.Now().Year())
			},
		},
	).ParseFiles(
		*tmpls...,
	))
}

func getFiles(path string, tmpls *[]string) *[]string {
	// Находим все что есть в папке
	matches, _ := filepath.Glob(path)
	// Перебираем элементы
	for _, el := range matches {
		info, _ := os.Stat(el)
		if !info.IsDir() {
			*tmpls = append(*tmpls, el)
		} else {
			getFiles(fmt.Sprintf("%s/*", el), tmpls)
		}
	}
	return tmpls
}
