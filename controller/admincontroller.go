package controller

import (
	//"html/template"

	"log"
	"net/http"
	"strconv"

	"github.com/sahasumit/BookWorm/model"
	"github.com/sahasumit/BookWorm/view"
)

//Admin Home page
/*
func Admin(res http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("HTMLS/admin/admin.html")
	t.Execute(res, nil)
}
*/

//list of  All unpublished book for admin
func UnPublishedBook(res http.ResponseWriter, req *http.Request) {
	log.Println(req.URL.Path)
	session, _ := store.Get(req, "cookie-name")
	userId := session.Values["UserID"].(int)
	userType := session.Values["UserType"].(string)

	log.Println("Admin looking for unpublished book = ", userId, userType)
	if userType != "admin" {
		session.Save(req, res)
		http.Redirect(res, req, "/user-home", 302)
		return
	}

	log.Println("Method :UnpublishedBook in Controller, List of All unpublished book and only admin can View")
	var data model.UData
	data.Books = model.GetBookList(0, 0)
	session.Save(req, res)
	view.UnPublishedBook(res, req, data)
}

//admin reviewing single book for publishing
func AdminReviewBook(res http.ResponseWriter, req *http.Request) {

	//	log.Println(userId, userType)
	session, _ := store.Get(req, "cookie-name")
	//	userId := session.Values["UserID"].(int)
	userType := session.Values["UserType"].(string)
	if userType != "admin" {
		session.Save(req, res)
		http.Redirect(res, req, "/user-home", 302)
		return
	}
	var data model.UData
	var book_id = req.URL.Query().Get("book")
	log.Println("Package : controller , Method : Admin review book, BookId ", book_id)
	bid, _ := strconv.Atoi(book_id)
	book := model.GetBook(bid)
	data.Book1 = book
	session.Save(req, res)
	view.AdminReviewBook(res, req, data)
}

func ApproveBook(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "cookie-name")

	userId := session.Values["UserID"].(int)
	userType := session.Values["UserType"].(string)
	log.Println(userId, userType)
	if userType != "admin" {
		session.Save(req, res)
		http.Redirect(res, req, "/user-home", 302)
		return
	}
	//
	var uid int

	uid = userId
	//
	book_id := req.URL.Query().Get("book")

	//	bid, _ = strconv.Atoi(book_id)
	bid, _ := strconv.Atoi(book_id)
	var data model.UData
	data.Book1 = model.GetBook(bid)
	log.Println("Package : Controller,Method : Approve Book, Book ID ; ", book_id, " PUBLISHER iD : ", data.Book1.PubId, " Logged In Id: ", uid)
	if data.Book1.PubId == uid {
		session.Save(req, res)
		http.Redirect(res, req, "/un-published-book", 302)
		return
	}
	//
	log.Println("Book to be approved is = " + book_id)

	model.PublishBook(bid, 1)
	session.Save(req, res)
	http.Redirect(res, req, "/un-published-book", 302)
}

func RejectBook(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "cookie-name")
	userId := session.Values["UserID"].(int)
	userType := session.Values["UserType"].(string)
	log.Println(userId, userType)
	if userType != "admin" {
		session.Save(req, res)
		http.Redirect(res, req, "/user-home", 302)
		return
	}
	book_id := req.URL.Query().Get("book")
	log.Println("Book to be rejected is = " + book_id)
	bid, _ := strconv.Atoi(book_id)
	model.PublishBook(bid, 2)
	session.Save(req, res)
	http.Redirect(res, req, "/un-published-book", 302)
}
func UnpublishBook(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "cookie-name")
	userId := session.Values["UserID"].(int)
	userType := session.Values["UserType"].(string)
	log.Println(userId, userType)
	if userType != "admin" {
		session.Save(req, res)
		http.Redirect(res, req, "/user-home", 302)
		return
	}
	book_id := req.URL.Query().Get("book")
	log.Println("Book to be unpublished is = " + book_id)
	bid, _ := strconv.Atoi(book_id)
	model.PublishBook(bid, 0)
	model.UnSubForAll(bid)
	session.Save(req, res)
	http.Redirect(res, req, "/publishedbook", 302)
}

func SendNotification(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "cookie-name")
	userId := session.Values["UserID"].(int)
	userType := session.Values["UserType"].(string)
	log.Println(userId, userType)
	if userType != "admin" {
		session.Save(req, res)
		http.Redirect(res, req, "/user-home", 302)
		return
	}
	bookId := req.URL.Query().Get("book")
	var data model.UData
	bid, _ := strconv.Atoi(bookId)
	data.Book1 = model.GetBook(bid)
	session.Save(req, res)
	view.SendNoti(res, req, data)
}

func PostNotification(res http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "cookie-name")
	userId := session.Values["UserID"].(int)
	userType := session.Values["UserType"].(string)
	log.Println(userId, userType)
	if userType != "admin" {
		session.Save(req, res)
		http.Redirect(res, req, "/user-home", 302)
		return
	}
	bookId := req.URL.Query().Get("book")
	bid, _ := strconv.Atoi(bookId)
	var nd model.Notification
	nd.BookId = bid
	nd.AdminNotification = req.FormValue("noti")
	model.SendNotification(nd)
	session.Save(req, res)
	http.Redirect(res, req, "/un-published-book", 302)
}

func UserList(res http.ResponseWriter, req *http.Request) {

	session, _ := store.Get(req, "cookie-name")
	userId := session.Values["UserID"].(int)
	userType := session.Values["UserType"].(string)
	log.Println(userId, userType)
	if userType != "admin" {
		session.Save(req, res)
		http.Redirect(res, req, "/user-home", 302)
		return
	}
	log.Println("Package : Controller ,Method : UserList  Admin ", userId, " Entered to view user list")
	var data model.UData
	//data.Users = model.GetUserList()
	session.Save(req, res)
	view.UserList(res, req, data)
}

func UserControl(res http.ResponseWriter, req *http.Request) {

	session, _ := store.Get(req, "cookie-name")
	userId := session.Values["UserID"].(int)
	userType := session.Values["UserType"].(string)
	log.Println(userId, userType)
	if userType != "admin" {
		session.Save(req, res)
		http.Redirect(res, req, "/user-home", 302)
		return
	}
	if req.Method != "Get" {
		return
	}
	userid := req.URL.Query().Get("userid")
	isBlock := req.URL.Query().Get("doblock")
	var uid, is int
	uid, _ = strconv.Atoi(userid)
	is, _ = strconv.Atoi(isBlock)
	var isb int
	if is == 0 {
		isb = 1
	} else {

		isb = 0
	}

	model.SetActiveUser(uid, isb)
	session.Save(req, res)
	http.Redirect(res, req, "/user-list", 302)
}
