package controller

import (
	"html/template"
	"net/http"
)

func Publisher(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("HTMLS/publisher/publisher.html")
	t.Execute(res, nil)
}
