package controller

import (
	"html"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"github.com/sahasumit/BookWorm/model"
	"github.com/sahasumit/BookWorm/view"
)

//----------------------
/*
func Test(res http.ResponseWriter, req *http.Request) {
	log.Println("Package : Controller, Method : test ")
	session, _ := store.Get(req, "cookie-name")
	log.Println("Logged In User Id : ", session.Values["UserId"])
	log.Println("Logged In User Type : ", session.Values["UserType"])
	for k := range session.Values {
		delete(session.Values, k)
	}
	log.Println("After Clearing Session")
	session, _ = store.Get(req, "cookie-name")
	log.Println("Logged In User Id : ", session.Values["UserId"])
	log.Println("Logged In User Type : ", session.Values["UserType"])

}
*/
//---------------------
func generateSessionID(email string, userID int)string{
	return email
}

type User struct {
	UserID  int
	UserType  string
	LoggedIn bool
}

func newPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

func setStruct(c redis.Conn, sessionID string, userID int, userType string) error {
	//const objectPrefix string = "user:"
	sessionInformation := User{
		UserID:  userID,
		UserType:userType,
		LoggedIn:true,
	}
	// serialize User object to JSON
	json, err := json.Marshal(sessionInformation)
	if err != nil {
		return err
	}

	// SET object
	_, err = c.Do("SET",sessionID, json)
	if err != nil {
		return err
	}
  log.Println(sessionID)
	return nil
}

func getStruct(c redis.Conn, sessionID string) User {
	s, err := redis.String(c.Do("GET", sessionID))
	if err == redis.ErrNil {
		log.Println("User does not exist")
	}
	sessionInformation := User{}
	err = json.Unmarshal([]byte(s), &sessionInformation)
//log.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	//log.Println(sessionInformation.UserType)
	return sessionInformation
}

func setSessionRedis(sessionID string, userID int, userType string){
	pool := newPool()
	conn := pool.Get()
	defer conn.Close()
	setStruct(conn, sessionID, userID, userType)
	//getStruct(conn, sessionID)
}

func Login(res http.ResponseWriter, req *http.Request) {
	clearSession(req)
	session, _ := store.Get(req, "cookie-name")
	var data model.UData
	log.Println("Logedin user = " + data.User1.Name)
	//processing GET method
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
