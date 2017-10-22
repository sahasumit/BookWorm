package view

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"model"
	"net/http"
	"strings"
)

var templates *template.Template
var err error

func Init() {
	var allFiles []string
	files, err := ioutil.ReadDir("./templates")
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, "./templates/"+filename)
		}
	}
	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		log.Println(err)
	}
}

func Home(res http.ResponseWriter, req *http.Request, data interface{}) {
	t := templates.Lookup("home.html")
	t.ExecuteTemplate(res, "home", data)
}

func Login(res http.ResponseWriter, req *http.Request, data model.UData) {
	t3 := templates.Lookup("login.html")
	t3.ExecuteTemplate(res, "login", data)
}

func UserHome(res http.ResponseWriter, req *http.Request, data model.UData) {
	t := templates.Lookup("user-home.html")
	if data.User1.UserType == "admin" {
		t.ExecuteTemplate(res, "admin-home", data)
	} else if data.User1.UserType == "publisher" {
		t.ExecuteTemplate(res, "publisher-home", data)
	} else if data.User1.UserType == "member" {
		t.ExecuteTemplate(res, "member-home", data)
	}
}
