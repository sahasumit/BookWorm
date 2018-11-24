package Api

import (
	"encoding/json"
	"net/http"

	"github.com/sahasumit/BookWorm/model"
)

func GetUsers(res http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		return
	}
	//log.Println("Hello GetUserList" + req.Method)
	var data model.UData
	data.Users = model.GetUserList()
	json.NewEncoder(res).Encode(&data.Users)
}

/*
func ReadBook(res http.ResponseWriter, req *http.Request) {
	//session, _ := store.Get(req, "cookie-name")
	//userId := session.Values["UserID"].(int)
	//userType := session.Values["UserType"].(string)
	//log.Println(userId, userType)
	//if userType != "admin" {
	//		session.Save(req, res)
	//	http.Redirect(res, req, "/login", 302)
	//	return
	//}
	var book_id = req.URL.Query().Get("book")
	//session.Save(req, res)
	http.Redirect(res, req, "/uploads/Pdf/"+book_id+".pdf", 302)
}
*/
