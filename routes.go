package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"html/template"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter(db DataHandler) *mux.Router {

	fe := FrontEnd{DataHandler: db}
	fe.CookieHandler = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

	var routes = Routes{
		Route{"Index", "GET", "/", Index},
	}

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./www/")))

	return router

}

type DataHandler interface {

}

type FrontEnd struct {
	DataHandler
	CookieHandler *securecookie.SecureCookie
}

type Page struct {
	PageData interface{}
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	var err error
	tmpl := template.New("")

	if tmpl, err = template.ParseFiles("templates/layout.tmpl", filename); err != nil {
		fmt.Println(err);
		return
	}

	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	render(w, "templates/home.tmpl", nil)
}

