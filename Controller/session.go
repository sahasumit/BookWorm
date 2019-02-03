package controller

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

//func getUser(req *http.Request) (userID, userType string) {
//	session, _ := store.Get(req, "cookie-name")
//	userID := session.Values["UserID"]
//	userType: = session.Values["UserType"]
//	return userID, userType
//}

func setSession(sessionID string, userID int, userType string, req *http.Request) {
	session, _ := store.Get(req, "cookie-name")
	session.Values["sessionID"] = sessionID
	//session.Values["UserID"] = userID
	//session.Values["UserType"] = userType
	//session.Values["LoggedIn"] = "True"
}

func clearSession(req *http.Request) {
	session, _ := store.Get(req, "cookie-name")

	for k := range session.Values {
		delete(session.Values, k)
	}
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
