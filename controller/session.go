package controller

import (
	//    "fmt"
	"net/http"

	"github.com/gorilla/securecookie"
)

// cookie handling

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

func getUser(req *http.Request) (userID, userType string) {
	cookie, err := req.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userID = cookieValue["name"]
			userType = cookieValue["userType"]
		}
	}
	return userID, userType
}

func setSession(userID string, userType string, res http.ResponseWriter) {
	value := map[string]string{
		"name":     userID,
		"userType": userType,
	}
	encoded, err := cookieHandler.Encode("session", value)
	if err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(res, cookie)
	}
}

func clearSession(res http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(res, cookie)
}

/*
// login handler

func loginHandler(response http.ResponseWriter, request *http.Request) {
    name := request.FormValue("name")
    pass := request.FormValue("password")
    redirectTarget := "/"
    if name != "" && pass != "" {
        // .. check credentials ..
        setSession(name, response)
        redirectTarget = "/internal"
    }
    http.Redirect(response, request, redirectTarget, 302)
}

// logout handler

func logoutHandler(response http.ResponseWriter, request *http.Request) {
    clearSession(response)
    http.Redirect(response, request, "/", 302)
}

// index page

const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`

func indexPageHandler(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, indexPage)
}

// internal page

const internalPage = `
<h1>Internal</h1>
<hr>
<small>User: %s</small>
<form method="post" action="/logout">
    <button type="submit">Logout</button>
</form>
`

func internalPageHandler(response http.ResponseWriter, request *http.Request) {
    userName := getUserName(request)
    if userName != "" {
        fmt.Fprintf(response, internalPage, userName)
    } else {
        http.Redirect(response, request, "/", 302)
    }
}

// server main method

var router = mux.NewRouter()

func main() {

    router.HandleFunc("/", indexPageHandler)
    router.HandleFunc("/internal", internalPageHandler)

    router.HandleFunc("/login", loginHandler).Methods("POST")
    router.HandleFunc("/logout", logoutHandler).Methods("POST")

    http.Handle("/", router)
    http.ListenAndServe(":9000", nil)
}

*/
