package controller

import (
	"html"
	"log"
	"net/http"
	"github.com/sahasumit/BookWorm/model"
	"github.com/sahasumit/BookWorm/view"
)

func generateSessionID(email string, userID int)string{
	return email
}

func Login(res http.ResponseWriter, req *http.Request) {
	clearSession(req)
	session, _ := store.Get(req, "cookie-name")
	var data model.UData
	log.Println("Logedin user = " + data.User1.Name)
	if req.Method != "POST" {

		log.Println("Serving login Page! ")
		session.Save(req, res)
		view.Login(res, req, data)
		return
	}

	//processing POST method
	req.ParseForm()
	email := html.EscapeString(req.FormValue("email"))
	password := html.EscapeString(req.FormValue("password"))
	log.Println("User Login Attempt by: ", email, " ", password)
	var user model.User
	user = model.GetUser(email)

	if user.Email != email {
		log.Println("User not found")
		data.Message = "Invalid Email!"
		session.Save(req, res)
		view.Login(res, req, data)
		return
	}
	if user.Password != password {
		log.Println("Password does not match")
		data.Message = "Incorrect Password!!"
		session.Save(req, res)
		view.Login(res, req, data)
		return
	}

	//if user is blocked redirect him
	if user.IsActive == 0 {
		log.Println("User is blocked")
		data.Message = "User is Blocked!"
		session.Save(req, res)
		view.Login(res, req, data)
		return
	}

  //set session id to redis
  sessionID := generateSessionID(user.Email, user.UserId)
	setSessionRedis(sessionID, user.UserId, user.UserType)

	//------------------------------------------

	//Set Session for newly loggedIn user here****
	//**********************************************
	//	LoggedInUser = user
	//uid := strconv.Itoa(user.UserId)

	//set Session
	//----------------------------
	setSession(sessionID, user.UserId, user.UserType, req)
	//----------------------------
	session.Save(req, res)
	log.Println("Welcome success login id = ", user.UserId, " Name = "+user.Name)
	//redirect according to user type
	data.Message = "Welcome " + user.Name + "! Login Succesful!!"
	data.User1 = user
	session.Save(req, res)
	http.Redirect(res, req, "/user-home", 302) // redirect to user home(admin/pub/member)
}
