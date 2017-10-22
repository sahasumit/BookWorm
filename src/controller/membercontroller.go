package controller

import (
	"html/template"
	"net/http"
)

func Member(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("HTMLS/member/member.html")
	t.Execute(res, nil)
}
