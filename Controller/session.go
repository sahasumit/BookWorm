package controller

import (
	"net/http"
	"github.com/gorilla/sessions"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
		"log"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)
//browser side section--------------------------------------------------------------
// set session data at browser cookie
func setSession(sessionID string, userID int, userType string, req *http.Request) {

	session, _ := store.Get(req, "cookie-name")
	session.Values["sessionID"] = sessionID
}

// remove session data from browser cookie
func clearSession(req *http.Request) {
	session, _ := store.Get(req, "cookie-name")

	for k := range session.Values {
		delete(session.Values, k)
	}
}
//------------------------------------------------------------------------------------------

// server side session data storing in redis------------------------------------------------
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
	return sessionInformation
}

func setSessionRedis(sessionID string, userID int, userType string){
	pool := newPool()
	conn := pool.Get()
	defer conn.Close()
	setStruct(conn, sessionID, userID, userType)
}
//-----------------------------------------------------------------------------------------------

//validate session data from redis------------------------------------------------------------------------------------------
func validateLogin(res http.ResponseWriter, req *http.Request) User{
	session, _ := store.Get(req, "cookie-name")
	sessionID := session.Values["sessionID"].(string)
	session.Save(req, res)
	pool := newPool()
	conn := pool.Get()
	defer conn.Close()
	sessionInformation := getStruct(conn, sessionID)
	return sessionInformation
}
//-----------------------------------------------------------------------------------------
