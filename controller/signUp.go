package controller

import (
	"html"
	"log"
	"net/http"

	"github.com/sahasumit/BookWorm/model"
	"github.com/sahasumit/BookWorm/view"
)

func SignUp(res http.ResponseWriter, req *http.Request) {

	clearSession(req)
	session, _ := store.Get(req, "cookie-name")

	var data model.UData
	log.Println("Entered Method : SignUp")
	//before clicking submit option
	if req.Method != "POST" {
		session.Save(req, res)
		view.SignUp(res, req, data)
		return
	}

	//getting signup information
	req.ParseForm()
	name := html.EscapeString(req.FormValue("name"))
	email := html.EscapeString(req.FormValue("email"))
	password1 := html.EscapeString(req.FormValue("password1"))
	password2 := html.EscapeString(req.FormValue("password2"))
	usertype := html.EscapeString(req.FormValue("UserType"))
	log.Println("Name ", name, "Email ", email, "password1 ", password1, "password2 ", password2, " Type ", usertype)

	//matching password for confirmation
	if password1 != password2 {
		log.Println("Password does not match")
		data.Message = "Password does not match"
		session.Save(req, res)
		view.SignUp(res, req, data)
		//http.Redirect(res, req, "/signup", 302)
		return
	}
	//checking mail used or not
	var emailexist string
	var user model.User
	user = model.GetUser(email)
	emailexist = user.Email
	if emailexist == email {
		log.Println("Email already used")
		data.Message = "Email already used"
		session.Save(req, res)
		view.SignUp(res, req, data)
		return
	}
	//generating unique user id
	var user_id int
	user_id = model.GenerateID(1)
	user.Set(user_id, email, password1, name, 1, usertype)
	model.SetUser(user)
	println("Sign Up successfull ", user_id)
	//storing new user in database user tab
	println("Stored in database")
	session.Save(req, res)
	http.Redirect(res, req, "/login", 302)
}
